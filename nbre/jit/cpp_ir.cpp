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
#include "jit/cpp_ir.h"
#include "fs/util.h"

namespace neb {
namespace cpp {
std::atomic_int cpp_ir::s_file_counter(1);

cpp_ir::cpp_ir(const std::string &cpp_content)
    : m_cpp_content(cpp_content), m_b_got_error(false) {}

neb::util::bytes cpp_ir::llvm_ir_content() {
  if (m_llvm_ir_fp == std::string("") && !m_b_got_error) {

    std::string fp_base = generate_fp();
    std::string cpp_fp = fp_base + ".cpp";
    std::string ir_fp = fp_base + ".bc";
    m_cpp_fp = cpp_fp;
    m_llvm_ir_fp = ir_fp;

    std::ofstream ofp(cpp_fp);
    ofp.write(m_cpp_content.c_str(), m_cpp_content.size());
    ofp.close();
    auto ret = make_ir_bitcode(cpp_fp, ir_fp);
    if (ret != 0) {
      m_b_got_error = true;
    }
  }
  if (m_b_got_error) {
    return neb::util::bytes();
  }
  if (!::neb::fs::exists(m_llvm_ir_fp)) {
    return neb::util::bytes();
  }

  std::ifstream ifs;
  ifs.open(m_llvm_ir_fp.c_str(), std::ios::in | std::ios::binary);
  if (!ifs.is_open())
    return neb::util::bytes();
  ifs.seekg(0, ifs.end);
  std::ifstream::pos_type size = ifs.tellg();
  neb::util::bytes buf(size);
  ifs.seekg(0, ifs.beg);
  ifs.read((char *)buf.value(), buf.size());
  ifs.close();
  return buf;
}
std::string cpp_ir::generate_fp() {
  std::string temp_path = neb::fs::tmp_dir();

  auto t = neb::fs::join_path(temp_path, std::to_string(s_file_counter));
  s_file_counter++;
  return t;
}

int cpp_ir::make_ir_bitcode(const std::string &cpp_file,
                            const std::string &ir_bc_file) {
  int result = -1;

  std::string current_path = neb::fs::cur_dir();
  std::string command_string(
      neb::fs::join_path(current_path, "lib_llvm/bin/clang") +
      " -O2 -emit-llvm -stdlib=libc++ -nodefaultlibs -fno-exceptions ");

  std::string include_string =
      std::string("-I") +
      neb::fs::join_path(::neb::configuration::instance().nbre_root_dir(),
                         "lib/include") +
      std::string(" ");

  command_string += include_string + " -o " + ir_bc_file;
  command_string += std::string(" ") + cpp_file;
  std::cout << command_string << std::endl;
  LOG(INFO) << command_string;

  result = command_executor::execute_command(command_string);
  if (result != 0) {
    LOG(ERROR) << "error: executed by boost::process::system.";
    LOG(ERROR) << "result code = " << result;
    return -1;
  }
  return 0;
}

} // namespace cpp
} // namespace neb