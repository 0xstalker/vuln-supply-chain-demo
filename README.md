# Vulnerable Supply Chain Demo

**WARNING: This project contains intentional security vulnerabilities and should NEVER be run in any production or trusted environment.**

## Purpose

This repository serves as an educational demonstration of common supply chain security vulnerabilities and insecure coding practices in Go applications. It is designed for security researchers, developers, and DevOps teams to understand and identify these risks.

## Vulnerabilities Demonstrated

### Application-Level Vulnerabilities:
1. **Command Injection** - The `/execute?cmd=` endpoint directly executes user-provided commands without sanitization
2. **Server-Side Request Forgery (SSRF)** - The `/fetch?url=` endpoint fetches arbitrary URLs without validation
3. **Environment Variable Disclosure** - The `/env` endpoint exposes all environment variables

### Supply Chain Vulnerabilities:
1. **Outdated Dependencies** - Uses older versions of packages that may contain known vulnerabilities
2. **Build-Time Data Exfiltration** - The Dockerfile contains commands that simulate sending sensitive build information to external services
3. **Insecure CI/CD Pipeline** - GitHub Actions workflow demonstrates multiple security misconfigurations:
   - Secrets exposed in logs
   - Installing tools from untrusted sources
   - Pushing images without proper security checks

## Setup (For Educational Purposes Only)

```bash
# Clone the repository
git clone https://github.com/0xstalker/vuln-supply-chain-demo.git
cd vuln-supply-chain-demo

# Install dependencies
go mod download

# Run the application
go run main.go
```

## Endpoints

- `GET /` - Home page with endpoint documentation
- `GET /execute?cmd=<command>` - Executes the specified command (VULNERABLE)
- `GET /fetch?url=<url>` - Fetches data from the specified URL (VULNERABLE)
- `GET /env` - Displays environment variables (POTENTIALLY VULNERABLE)

## Security Disclaimer

**DO NOT RUN THIS APPLICATION IN ANY PRODUCTION ENVIRONMENT OR SYSTEM YOU CARE ABOUT.**

This code contains multiple severe security vulnerabilities deliberately introduced for educational purposes. Running this code could result in:
- Remote code execution
- Data exfiltration
- Unauthorized system access
- Information disclosure

## Educational Objectives

This project demonstrates:
1. Importance of input validation and sanitization
2. Risks of executing user-controlled commands
3. Dangers of allowing arbitrary network requests
4. Proper handling of environment variables and secrets
5. Secure dependency management practices
6. Safe CI/CD pipeline configurations

## License

This project is intended for educational use only. Use responsibly for learning and testing purposes.