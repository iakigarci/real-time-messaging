const jwt = require("jsonwebtoken");
const winston = require("winston");
const User = require("./models");

class AuthService {
  constructor(jwtSecret) {
    this.jwtSecret = jwtSecret;
    this.logger = winston.createLogger({
      level: "info",
      format: winston.format.json(),
      transports: [
        new winston.transports.Console(),
        new winston.transports.File({ filename: "error.log", level: "error" }),
        new winston.transports.File({ filename: "combined.log" }),
      ],
    });
  }

  async generateToken(user) {
    if (!(user instanceof User)) {
      throw new Error("Invalid user object");
    }

    try {
      const claims = {
        user_id: user.email,
        email: user.email,
        exp: Math.floor(Date.now() / 1000) + 24 * 60 * 60, // 24 hours from now
        iat: Math.floor(Date.now() / 1000),
      };

      const token = jwt.sign(claims, this.jwtSecret, { algorithm: "HS256" });
      return token;
    } catch (error) {
      this.logger.error("Failed to generate token:", error);
      throw error;
    }
  }

  async validateToken(tokenString) {
    try {
      const decoded = jwt.verify(tokenString, this.jwtSecret, {
        algorithms: ["HS256"],
      });

      if (!decoded.user_id) {
        const error = new Error("Invalid user id");
        this.logger.error(error.message);
        throw error;
      }

      return decoded.user_id;
    } catch (error) {
      this.logger.error("Failed to validate token:", error);
      throw error;
    }
  }
}

module.exports = AuthService;
