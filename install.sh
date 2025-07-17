#!/bin/bash
# LupettoGo Installation Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🐺 LupettoGo Installation Script${NC}"
echo "============================================"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="x86_64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}❌ Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

case $OS in
    linux)
        OS="linux"
        ;;
    darwin)
        OS="darwin"
        ;;
    *)
        echo -e "${RED}❌ Unsupported OS: $OS${NC}"
        exit 1
        ;;
esac

echo -e "${YELLOW}Detected OS: $OS${NC}"
echo -e "${YELLOW}Detected Architecture: $ARCH${NC}"
echo

# Get latest release version
echo -e "${BLUE}📡 Fetching latest release...${NC}"
LATEST_VERSION=$(curl -s https://api.github.com/repos/adipras/lupettogo/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}❌ Failed to fetch latest version${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Latest version: $LATEST_VERSION${NC}"

# Download URL
DOWNLOAD_URL="https://github.com/adipras/lupettogo/releases/download/$LATEST_VERSION/lupettogo_${OS^}_${ARCH}.tar.gz"

# Download and install
echo -e "${BLUE}📥 Downloading LupettoGo...${NC}"
curl -L "$DOWNLOAD_URL" -o "/tmp/lupettogo.tar.gz"

echo -e "${BLUE}📦 Extracting...${NC}"
tar -xzf "/tmp/lupettogo.tar.gz" -C "/tmp"

echo -e "${BLUE}🔧 Installing to /usr/local/bin...${NC}"
sudo mv "/tmp/lupettogo" "/usr/local/bin/lupettogo"
sudo chmod +x "/usr/local/bin/lupettogo"

# Cleanup
rm -f "/tmp/lupettogo.tar.gz"

echo
echo -e "${GREEN}✅ LupettoGo installed successfully!${NC}"
echo
echo -e "${BLUE}🚀 Quick start:${NC}"
echo -e "  ${YELLOW}lupettogo init my-saas-app${NC}"
echo -e "  ${YELLOW}cd my-saas-app${NC}"
echo -e "  ${YELLOW}go mod tidy && go run main.go${NC}"
echo
echo -e "${BLUE}📖 For more information:${NC}"
echo -e "  ${YELLOW}lupettogo --help${NC}"
echo -e "  ${YELLOW}lupettogo doctor${NC}"
echo
echo -e "${GREEN}🐺 With the little wolf, no project is too big!${NC}"