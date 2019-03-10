// Copyright (C) 2018 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or
// modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see
// <http://www.gnu.org/licenses/>.
//
#include "core/net_ipc/client/nipc_client.h"
#include "common/configuration.h"
#include "core/net_ipc/nipc_pkg.h"

namespace neb {
namespace core {

nipc_client::~nipc_client() { shutdown(); }

bool nipc_client::start() {

  if (m_handlers.empty()) {
    LOG(INFO) << "No handlers here";
    return false;
  }

  bool init_done = false;
  std::mutex local_mutex;
  std::condition_variable local_cond_var;
  std::atomic_bool got_exception_when_start_ipc;

  m_thread = std::unique_ptr<std::thread>(new std::thread([&, this]() {
    try {
      got_exception_when_start_ipc = false;

      ::ff::net::net_nervure nn;
      ::ff::net::typed_pkg_hub hub;
      std::for_each(
          m_handlers.begin(), m_handlers.end(),
          [&hub](const std::function<void(::ff::net::typed_pkg_hub &)> &f) {
            f(hub);
          });

      m_to_recv_heart_beat_msg = 0;

      hub.to_recv_pkg<heart_beat_t>([this](std::shared_ptr<heart_beat_t>) {
        m_to_recv_heart_beat_msg--;
      });

      nn.get_event_handler()->listen<::ff::net::event::tcp_get_connection>(
          [&, this](::ff::net::tcp_connection_base *) {
            LOG(INFO) << "got connection";
            m_is_connected = true;
            local_mutex.lock();
            init_done = true;
            local_mutex.unlock();
            local_cond_var.notify_one();
          });
      nn.get_event_handler()->listen<::ff::net::event::tcp_lost_connection>(
          [this](::ff::net::tcp_connection_base *) {
            LOG(INFO) << "lost connection";
            m_is_connected = false;
            exit(-2);
          });
      nn.add_pkg_hub(hub);
      m_conn = nn.add_tcp_client(configuration::instance().nipc_listen(),
                                 configuration::instance().nipc_port());

      m_heart_bear_timer = std::make_unique<util::timer_loop>(&nn.ioservice());
      m_heart_bear_timer->register_timer_and_callback(3, [this]() {
        if (m_to_recv_heart_beat_msg > 2) {
          m_conn->close();
          command_queue::instance().send_command(
              std::make_shared<exit_command>());
        }
        m_to_recv_heart_beat_msg++;
        std::shared_ptr<heart_beat_t> hb = std::make_shared<heart_beat_t>();
        m_conn->send(hb);
      });
      nn.run();

    } catch (const std::exception &e) {
      got_exception_when_start_ipc = true;
      LOG(ERROR) << "get exception when start ipc, " << typeid(e).name() << ", "
                 << e.what();
      local_cond_var.notify_one();
    } catch (...) {
      got_exception_when_start_ipc = true;
      LOG(ERROR) << "get unknown exception when start ipc";
      local_cond_var.notify_one();
    }
  }));
  std::unique_lock<std::mutex> _l(local_mutex);
  if (!init_done) {
    local_cond_var.wait(_l);
  }
  if (got_exception_when_start_ipc) {
    return false;
  }
  return true;
}

void nipc_client::shutdown() {
  m_conn->close();
  if (m_thread) {
    m_thread->join();
    m_thread.reset();
  }
}
} // namespace core
} // namespace neb
