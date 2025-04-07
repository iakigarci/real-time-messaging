const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");
const AuthService = require("./authService");
const User = require("./models");

// Load proto file
const PROTO_PATH = path.join(__dirname, "..", "proto/auth.proto");
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const authProto = grpc.loadPackageDefinition(packageDefinition).auth;

const jwtSecret = process.env.JWT_SECRET || "secret";
const authService = new AuthService(jwtSecret);

const serviceImplementation = {
  generateToken: async (call, callback) => {
    try {
      const { email } = call.request;
      const user = new User(email);
      const token = await authService.generateToken(user);
      callback(null, { token });
    } catch (error) {
      callback({
        code: grpc.status.INTERNAL,
        details: `Error generating token: ${error.message}`,
      });
    }
  },

  validateToken: async (call, callback) => {
    try {
      const { token } = call.request;
      const userId = await authService.validateToken(token);
      callback(null, { user_id: userId });
    } catch (error) {
      callback({
        code: grpc.status.INTERNAL,
        details: `Error validating token: ${error.message}`,
      });
    }
  },
};

// Create and start gRPC server
function startServer() {
  const server = new grpc.Server();
  server.addService(authProto.AuthService.service, serviceImplementation);

  const port = process.env.PORT || 50051;
  server.bindAsync(
    `0.0.0.0:${port}`,
    grpc.ServerCredentials.createInsecure(),
    (error, port) => {
      if (error) {
        console.error("Failed to bind server:", error);
        return;
      }
      server.start();
      console.log(`gRPC server running on port ${port}`);
    }
  );
}

startServer();
