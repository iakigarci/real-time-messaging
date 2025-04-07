const AuthService = require("../src/authService");
const User = require("../src/models");
const jwt = require("jsonwebtoken");

describe("AuthService", () => {
  const jwtSecret = "test-secret";
  let authService;
  let mockUser;

  beforeEach(() => {
    authService = new AuthService(jwtSecret);
    mockUser = new User();
    mockUser.email = "test@example.com";
    jest.resetAllMocks();
  });

  describe("generateToken", () => {
    it("should generate a valid JWT token for a valid user", async () => {
      const token = await authService.generateToken(mockUser);

      expect(token).toBeDefined();
      const decoded = jwt.verify(token, jwtSecret);
      expect(decoded.user_id).toBe(mockUser.email);
      expect(decoded.email).toBe(mockUser.email);
      expect(decoded.exp).toBeDefined();
      expect(decoded.iat).toBeDefined();
    });

    it("should throw error when invalid user object is provided", async () => {
      const invalidUser = { email: "test@example.com" };

      await expect(authService.generateToken(invalidUser)).rejects.toThrow(
        "Invalid user object"
      );
    });

    it("should throw error when JWT signing fails", async () => {
      jest.spyOn(jwt, "sign").mockImplementation(() => {
        throw new Error("Signing failed");
      });

      await expect(authService.generateToken(mockUser)).rejects.toThrow(
        "Signing failed"
      );

      // Reset the mock after this test
      jest.restoreAllMocks();
    });
  });

  describe("validateToken", () => {
    let validToken;

    beforeEach(async () => {
      jest.restoreAllMocks();
      validToken = await authService.generateToken(mockUser);
    });

    it("should successfully validate a valid token", async () => {
      const userId = await authService.validateToken(validToken);
      expect(userId).toBe(mockUser.email);
    });

    it("should throw error for invalid token", async () => {
      const invalidToken = "invalid.token.string";

      await expect(authService.validateToken(invalidToken)).rejects.toThrow();
    });

    it("should throw error for token without user_id", async () => {
      const tokenWithoutUserId = jwt.sign(
        { email: "test@example.com" },
        jwtSecret,
        { algorithm: "HS256" }
      );

      await expect(
        authService.validateToken(tokenWithoutUserId)
      ).rejects.toThrow("Invalid user id");
    });

    it("should throw error for expired token", async () => {
      const expiredToken = jwt.sign(
        {
          user_id: mockUser.email,
          email: mockUser.email,
          exp: Math.floor(Date.now() / 1000) - 3600, // 1 hour ago
        },
        jwtSecret,
        { algorithm: "HS256" }
      );

      await expect(authService.validateToken(expiredToken)).rejects.toThrow();
    });
  });
});
