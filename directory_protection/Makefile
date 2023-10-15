# Makefile для компиляции, запуска и создания архива tar.gz

# Имя выходного исполняемого файла
BINARY_NAME=directory_protection.exe

# Компилятор Go
GO=go

# Опции для компиляции
BUILD_FLAGS=-v

# Имя архива tar.gz
TGZ_ARCHIVE_NAME=bagautdinovlab1a.tar.gz

all: build

build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_NAME) main.go

run:
	$(GO) run main.go

archive-gzip: build
	tar czvf $(TGZ_ARCHIVE_NAME) $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME) $(TGZ_ARCHIVE_NAME)

.PHONY: all build run archive-gzip clean
