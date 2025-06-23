#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>
#include "generated/service.pb.h"
#include "generated/service.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using v1::EchoRequest;
using v1::EchoResponse;
using v1::EchoService;

class EchoServiceImpl final : public EchoService::Service {
  Status Echo(ServerContext* ctx, const EchoRequest* req,
              EchoResponse* resp) override {
    resp->set_message(req->message());
    return Status::OK;
  }
};

int main(int argc, char** argv) {
  std::string addr = "0.0.0.0:50051";
  if (argc > 1) addr = argv[1];

  EchoServiceImpl service;
  ServerBuilder builder;
  builder.AddListeningPort(addr, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<Server> server(builder.BuildAndStart());
  std::cout << "🚀 gRPC server listening on " << addr << std::endl;
  server->Wait();
  return 0;
}
