cmake_minimum_required(VERSION 3.25)
project({{ cookiecutter.project_slug }} LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Protobuf    REQUIRED)
find_package(gRPC        CONFIG REQUIRED)
find_package(Threads     REQUIRED)

# ── Proto generation ────────────────────────────────────────────────────────
set(PROTO_DIR   ${CMAKE_SOURCE_DIR}/../../proto/v1)
set(GEN_DIR     ${CMAKE_BINARY_DIR}/generated)
file(MAKE_DIRECTORY ${GEN_DIR})

find_program(PROTOC_EXE protoc      REQUIRED)
find_program(GRPC_PLUGIN grpc_cpp_plugin REQUIRED)

file(GLOB PROTO_FILES "${PROTO_DIR}/*.proto")
foreach(P ${PROTO_FILES})
  get_filename_component(NAME_WE ${P} NAME_WE)
  set(pb_cc  ${GEN_DIR}/${NAME_WE}.pb.cc)
  set(pb_h   ${GEN_DIR}/${NAME_WE}.pb.h)
  set(grcc   ${GEN_DIR}/${NAME_WE}.grpc.pb.cc)
  set(grch   ${GEN_DIR}/${NAME_WE}.grpc.pb.h)

  add_custom_command(
    OUTPUT ${pb_cc} ${pb_h} ${grcc} ${grch}
    COMMAND ${PROTOC_EXE}
            --cpp_out=${GEN_DIR}
            --grpc_out=${GEN_DIR}
            -I ${PROTO_DIR}
            --plugin=protoc-gen-grpc=${GRPC_PLUGIN}
            ${P}
    DEPENDS ${P}
  )
  list(APPEND PROTO_SRCS ${pb_cc} ${grcc})
  list(APPEND PROTO_HDRS ${pb_h}  ${grch})
endforeach()

# ── Server executable ───────────────────────────────────────────────────────
file(GLOB_RECURSE APP_SRCS ${CMAKE_SOURCE_DIR}/*.cc)
add_executable(${PROJECT} ${APP_SRCS} ${PROTO_SRCS} ${PROTO_HDRS})

target_include_directories(${PROJECT}
  PRIVATE
    ${GEN_DIR}
    ${PROTO_DIR}
)

target_link_libraries(${PROJECT}
  PRIVATE
    gRPC::grpc++ protobuf::libprotobuf Threads::Threads
)
