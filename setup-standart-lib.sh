#!/bin/bash

# Script untuk membuat struktur belajar Golang Standard Library
# Dibuat dengan urutan yang logis untuk pembelajaran

echo "🚀 Membuat struktur belajar Golang Standard Library..."

# Buat direktori utama
mkdir -p go-stdlib-learning
cd go-stdlib-learning

# 1. BASIC FUNDAMENTALS (Wajib dikuasai pertama)
echo "📁 Creating Basic Fundamentals..."
mkdir -p 01-basic-fundamentals/{fmt,strings,strconv,errors}

# fmt - Input/Output dasar
touch 01-basic-fundamentals/fmt/{printf.go,sprintf.go,scanf.go,print-variants.go,custom-formatting.go}
cat > 01-basic-fundamentals/fmt/README.md << 'EOF'
# Package fmt
Paket untuk formatted I/O operations.

## Yang akan dipelajari:
- Printf, Sprintf, Fprintf
- Print, Println variants  
- Scanf untuk input
- Custom formatting verbs
- Stringer interface
EOF

# strings - Manipulasi string
touch 01-basic-fundamentals/strings/{basic-operations.go,searching.go,replacing.go,splitting.go,case-conversion.go,builder.go}
cat > 01-basic-fundamentals/strings/README.md << 'EOF'
# Package strings
Operasi manipulasi string yang essential.

## Yang akan dipelajari:
- Contains, HasPrefix, HasSuffix
- Replace, ReplaceAll
- Split, Join
- ToUpper, ToLower, Title
- strings.Builder untuk performa
EOF

# strconv - Konversi tipe data
touch 01-basic-fundamentals/strconv/{string-to-number.go,number-to-string.go,boolean-conversion.go,advanced-parsing.go}
cat > 01-basic-fundamentals/strconv/README.md << 'EOF'
# Package strconv
Konversi antara string dan tipe data lain.

## Yang akan dipelajari:
- Atoi, ParseInt, ParseFloat
- Itoa, FormatInt, FormatFloat
- ParseBool, FormatBool
- Quote, Unquote
EOF

# errors - Error handling
touch 01-basic-fundamentals/errors/{basic-errors.go,custom-errors.go,wrapping-errors.go,error-types.go}
cat > 01-basic-fundamentals/errors/README.md << 'EOF'
# Package errors
Error handling yang proper di Go.

## Yang akan dipelajari:
- errors.New()
- Custom error types
- Error wrapping (Go 1.13+)
- errors.Is, errors.As
EOF

# 2. FILE & I/O OPERATIONS
echo "📁 Creating File & I/O Operations..."
mkdir -p 02-file-io/{io,os,bufio,path,filepath}

# io - Basic I/O operations
touch 02-file-io/io/{readers-writers.go,copy-operations.go,pipes.go,multiwriter.go,interfaces.go}
cat > 02-file-io/io/README.md << 'EOF'
# Package io
Interface dan fungsi I/O fundamental.

## Yang akan dipelajari:
- Reader, Writer interfaces
- io.Copy, io.CopyN
- io.Pipe untuk concurrent I/O
- MultiReader, MultiWriter
EOF

# os - Operating system interface
touch 02-file-io/os/{file-operations.go,directory-operations.go,environment.go,process.go,signals.go}
cat > 02-file-io/os/README.md << 'EOF'
# Package os
Interface ke operating system.

## Yang akan dipelajari:
- File operations (Create, Open, Remove)
- Directory operations (Mkdir, ReadDir)
- Environment variables
- Process operations
- Signal handling
EOF

# bufio - Buffered I/O
touch 02-file-io/bufio/{buffered-reader.go,buffered-writer.go,scanner.go,line-reading.go}
cat > 02-file-io/bufio/README.md << 'EOF'
# Package bufio
Buffered I/O untuk performa yang lebih baik.

## Yang akan dipelajari:
- bufio.Reader, bufio.Writer
- bufio.Scanner untuk parsing
- ReadLine, ReadBytes
- Flush operations
EOF

# path & filepath - Path manipulation
touch 02-file-io/path/{url-paths.go,basic-operations.go}
touch 02-file-io/filepath/{file-paths.go,walking-directories.go,pattern-matching.go,cross-platform.go}
cat > 02-file-io/path/README.md << 'EOF'
# Package path
Manipulasi URL paths.
EOF

cat > 02-file-io/filepath/README.md << 'EOF'
# Package filepath
Manipulasi file paths yang cross-platform.

## Yang akan dipelajari:
- Join, Split, Dir, Base
- Walk, WalkDir
- Match, Glob
- Abs, Rel
EOF

# 3. DATA STRUCTURES & ALGORITHMS
echo "📁 Creating Data Structures & Algorithms..."
mkdir -p 03-data-algo/{sort,container}

# sort - Sorting algorithms
touch 03-data-algo/sort/{basic-sorting.go,custom-sorting.go,search-operations.go,performance.go}
cat > 03-data-algo/sort/README.md << 'EOF'
# Package sort
Sorting dan searching algorithms.

## Yang akan dipelajari:
- sort.Ints, sort.Strings, sort.Float64s
- sort.Slice dengan custom comparison
- sort.Search untuk binary search
- Interface sort.Interface
EOF

# container - Data structures
mkdir -p 03-data-algo/container/{heap,list,ring}
touch 03-data-algo/container/heap/{priority-queue.go,heap-operations.go}
touch 03-data-algo/container/list/{doubly-linked-list.go,list-operations.go}
touch 03-data-algo/container/ring/{circular-list.go,ring-operations.go}

cat > 03-data-algo/container/README.md << 'EOF'
# Package container
Built-in data structures.

## Subpackages:
- heap: Priority queue implementation
- list: Doubly linked list
- ring: Circular list
EOF

# 4. TIME & DATE
echo "📁 Creating Time & Date..."
mkdir -p 04-time/time

touch 04-time/time/basic-time.go
touch 04-time/time/formatting-parsing.go
touch 04-time/time/duration.go
touch 04-time/time/timers.go
touch 04-time/time/zones.go
touch 04-time/time/ticker.go

cat > 04-time/time/README.md << 'EOF'
# Package time
Time dan date operations.

## Yang akan dipelajari:
- time.Now(), time.Parse(), time.Format()
- Duration calculations
- Timer dan Ticker
- Timezone handling
- Sleep operations
EOF

# 5. NETWORKING & HTTP
echo "📁 Creating Networking & HTTP..."
mkdir -p 05-networking/{net,http,url}

# net - Network programming
mkdir -p 05-networking/net/{tcp,udp,unix}
touch 05-networking/net/{basic-networking.go,listeners.go,connections.go}
touch 05-networking/net/tcp/{tcp-server.go,tcp-client.go}
touch 05-networking/net/udp/{udp-server.go,udp-client.go}
touch 05-networking/net/unix/{unix-sockets.go}

cat > 05-networking/net/README.md << 'EOF'
# Package net
Network programming primitives.

## Yang akan dipelajari:
- TCP/UDP connections
- Listeners dan Dialers
- Unix sockets
- IP address handling
EOF

# http - HTTP client/server
mkdir -p 05-networking/http/{client,server,middleware}
touch 05-networking/http/{basic-http.go,http-methods.go,headers.go,cookies.go}
touch 05-networking/http/client/{get-post.go,custom-client.go,timeouts.go}
touch 05-networking/http/server/{basic-server.go,routing.go,static-files.go}
touch 05-networking/http/middleware/{logging.go,auth.go,cors.go}

cat > 05-networking/http/README.md << 'EOF'
# Package net/http
HTTP client dan server implementation.

## Yang akan dipelajari:
- HTTP Client (GET, POST, custom headers)
- HTTP Server (routing, middleware)
- Request/Response handling
- File server
- WebSocket basics
EOF

# url - URL parsing
touch 05-networking/url/{url-parsing.go,query-params.go,url-building.go}
cat > 05-networking/url/README.md << 'EOF'
# Package net/url
URL parsing dan manipulation.

## Yang akan dipelajari:
- URL parsing dan validation
- Query parameters
- URL escaping/unescaping
EOF

# 6. CONCURRENCY
echo "📁 Creating Concurrency..."
mkdir -p 06-concurrency/{sync,context,runtime}

# sync - Synchronization primitives
touch 06-concurrency/sync/{mutex.go,rwmutex.go,waitgroup.go,once.go,cond.go,pool.go,map.go}
cat > 06-concurrency/sync/README.md << 'EOF'
# Package sync
Synchronization primitives untuk concurrency.

## Yang akan dipelajari:
- Mutex, RWMutex
- WaitGroup untuk goroutine coordination
- sync.Once untuk one-time initialization
- sync.Pool untuk object reuse
- sync.Map untuk concurrent map
EOF

# context - Context handling
touch 06-concurrency/context/{basic-context.go,cancellation.go,timeouts.go,values.go,request-context.go}
cat > 06-concurrency/context/README.md << 'EOF'
# Package context
Context untuk cancellation, timeouts, dan values.

## Yang akan dipelajari:
- context.Background(), context.TODO()
- WithCancel, WithTimeout, WithDeadline
- WithValue untuk request-scoped data
- Context best practices
EOF

# runtime - Runtime operations
touch 06-concurrency/runtime/{goroutines.go,gc.go,memory.go,cpu.go,debug.go}
cat > 06-concurrency/runtime/README.md << 'EOF'
# Package runtime
Interface ke Go runtime system.

## Yang akan dipelajari:
- GOMAXPROCS, NumGoroutine
- GC operations
- Memory statistics
- Stack traces
EOF

# 7. ENCODING & SERIALIZATION
echo "📁 Creating Encoding & Serialization..."
mkdir -p 07-encoding/{json,xml,base64,gob,csv}

# json - JSON operations
touch 07-encoding/json/{marshal-unmarshal.go,struct-tags.go,streaming.go,custom-json.go,validation.go}
cat > 07-encoding/json/README.md << 'EOF'
# Package encoding/json
JSON encoding dan decoding.

## Yang akan dipelajari:
- Marshal/Unmarshal
- Struct tags untuk field mapping
- Streaming dengan Encoder/Decoder
- Custom JSON marshaling
- JSON validation
EOF

# xml - XML operations
touch 07-encoding/xml/{marshal-unmarshal.go,struct-tags.go,streaming.go}
cat > 07-encoding/xml/README.md << 'EOF'
# Package encoding/xml
XML encoding dan decoding.
EOF

# base64 - Base64 encoding
touch 07-encoding/base64/{standard-encoding.go,url-encoding.go,custom-encoding.go}
cat > 07-encoding/base64/README.md << 'EOF'
# Package encoding/base64
Base64 encoding dan decoding.
EOF

# gob - Go binary format
touch 07-encoding/gob/{basic-gob.go,network-gob.go,interface-gob.go}
cat > 07-encoding/gob/README.md << 'EOF'
# Package encoding/gob
Go binary format untuk serialization.
EOF

# csv - CSV operations
touch 07-encoding/csv/{reading-csv.go,writing-csv.go,custom-delimiter.go}
cat > 07-encoding/csv/README.md << 'EOF'
# Package encoding/csv
CSV file processing.
EOF

# 8. CRYPTOGRAPHY & SECURITY
echo "📁 Creating Cryptography & Security..."
mkdir -p 08-crypto/{crypto,hash,tls}

# crypto - Basic cryptography
mkdir -p 08-crypto/crypto/{rand,aes,rsa,sha}
touch 08-crypto/crypto/{basic-crypto.go,symmetric.go,asymmetric.go}
touch 08-crypto/crypto/rand/{random-generation.go,secure-random.go}
touch 08-crypto/crypto/aes/{aes-encryption.go,gcm-mode.go}
touch 08-crypto/crypto/rsa/{rsa-keys.go,signing.go}
touch 08-crypto/crypto/sha/{hashing.go,hmac.go}

cat > 08-crypto/crypto/README.md << 'EOF'
# Package crypto
Cryptographic operations.

## Yang akan dipelajari:
- Random number generation
- Symmetric encryption (AES)
- Asymmetric encryption (RSA)
- Digital signatures
- Hashing (SHA family)
EOF

# hash - Hashing algorithms
touch 08-crypto/hash/{basic-hashing.go,checksums.go,custom-hash.go}
cat > 08-crypto/hash/README.md << 'EOF'
# Package hash
Hashing interfaces dan implementations.
EOF

# tls - TLS/SSL
touch 08-crypto/tls/{tls-client.go,tls-server.go,certificates.go}
cat > 08-crypto/tls/README.md << 'EOF'
# Package crypto/tls
TLS/SSL implementations.
EOF

# 9. REFLECTION & RUNTIME
echo "📁 Creating Reflection & Runtime..."
mkdir -p 09-reflection/{reflect,unsafe}

# reflect - Reflection
touch 09-reflection/reflect/{basic-reflection.go,types.go,values.go,struct-inspection.go,method-calling.go}
cat > 09-reflection/reflect/README.md << 'EOF'
# Package reflect
Runtime reflection capabilities.

## Yang akan dipelajari:
- reflect.Type, reflect.Value
- Struct field inspection
- Method calling via reflection
- Interface{} inspection
- Performance considerations
EOF

# unsafe - Unsafe operations
touch 09-reflection/unsafe/{pointer-operations.go,memory-layout.go,type-conversion.go}
cat > 09-reflection/unsafe/README.md << 'EOF'
# Package unsafe
Low-level memory operations (use with caution).

## Yang akan dipelajari:
- Pointer arithmetic
- Memory layout inspection
- Unsafe type conversions
- Performance optimizations
EOF

# 10. TESTING & BENCHMARKING
echo "📁 Creating Testing & Benchmarking..."
mkdir -p 10-testing/{testing,debug}

# testing - Testing framework
touch 10-testing/testing/{basic-tests.go,table-tests.go,benchmarks.go,examples.go,mocking.go,coverage.go}
cat > 10-testing/testing/README.md << 'EOF'
# Package testing
Testing dan benchmarking framework.

## Yang akan dipelajari:
- Basic unit tests
- Table-driven tests
- Benchmarking
- Example tests
- Test coverage
- Mocking strategies
EOF

# debug - Debugging utilities
mkdir -p 10-testing/debug/{pprof,trace}
touch 10-testing/debug/{stack-traces.go,memory-debug.go}
touch 10-testing/debug/pprof/{cpu-profiling.go,memory-profiling.go,web-interface.go}
touch 10-testing/debug/trace/{execution-tracing.go}

cat > 10-testing/debug/README.md << 'EOF'
# Package runtime/debug
Debugging dan profiling utilities.
EOF

# 11. ADVANCED TOPICS
echo "📁 Creating Advanced Topics..."
mkdir -p 11-advanced/{regexp,template,plugin,build}

# regexp - Regular expressions
touch 11-advanced/regexp/{basic-regexp.go,compiling.go,finding.go,replacing.go,groups.go}
cat > 11-advanced/regexp/README.md << 'EOF'
# Package regexp
Regular expression operations.

## Yang akan dipelajari:
- Regexp compilation
- Finding matches
- Replacing text
- Capture groups
- Performance tips
EOF

# template - Templating
mkdir -p 11-advanced/template/{text,html}
touch 11-advanced/template/text/{basic-templates.go,data-binding.go,functions.go}
touch 11-advanced/template/html/{html-templates.go,security.go,inheritance.go}

cat > 11-advanced/template/README.md << 'EOF'
# Package text/template & html/template
Template engines untuk text dan HTML.

## Yang akan dipelajari:
- Template syntax
- Data binding
- Custom functions
- Template inheritance
- XSS protection (html/template)
EOF

# plugin - Plugin system
touch 11-advanced/plugin/{loading-plugins.go,plugin-interface.go}
cat > 11-advanced/plugin/README.md << 'EOF'
# Package plugin
Dynamic plugin loading (Linux/macOS only).
EOF

# build - Build constraints
touch 11-advanced/build/{build-tags.go,conditional-compilation.go}
cat > 11-advanced/build/README.md << 'EOF'
# Build Constraints
Conditional compilation dengan build tags.
EOF

# Buat file README utama
cat > README.md << 'EOF'
# Golang Standard Library Learning Path 🚀

Struktur pembelajaran Golang Standard Library yang terorganisir dari basic hingga advanced.

## 📚 Learning Path

### 1. Basic Fundamentals (WAJIB PERTAMA)
- `fmt` - Input/Output formatting
- `strings` - String manipulation  
- `strconv` - String conversion
- `errors` - Error handling

### 2. File & I/O Operations
- `io` - I/O primitives
- `os` - Operating system interface
- `bufio` - Buffered I/O
- `path/filepath` - File path manipulation

### 3. Data Structures & Algorithms
- `sort` - Sorting dan searching
- `container/*` - Built-in data structures

### 4. Time & Date
- `time` - Time operations

### 5. Networking & HTTP
- `net` - Network programming
- `net/http` - HTTP client/server
- `net/url` - URL operations

### 6. Concurrency (SANGAT PENTING)
- `sync` - Synchronization primitives
- `context` - Context handling
- `runtime` - Runtime operations

### 7. Encoding & Serialization
- `encoding/json` - JSON operations
- `encoding/xml` - XML operations
- `encoding/csv` - CSV operations
- Other encoding formats

### 8. Cryptography & Security
- `crypto/*` - Cryptographic operations
- `hash` - Hashing algorithms

### 9. Reflection & Runtime
- `reflect` - Runtime reflection
- `unsafe` - Unsafe operations

### 10. Testing & Benchmarking
- `testing` - Testing framework
- `runtime/debug` - Debugging utilities

### 11. Advanced Topics
- `regexp` - Regular expressions
- `text/template` - Template engine
- Advanced build techniques

## 🎯 Best Practices

1. **Start with basics** - Mulai dari folder 01-basic-fundamentals
2. **Practice each concept** - Tulis code di setiap file yang disediakan
3. **Read the docs** - Selalu baca dokumentasi resmi Go
4. **Write tests** - Practice testing untuk setiap concept
5. **Build projects** - Combine multiple packages dalam project nyata

## 🚀 How to Use

1. Mulai dari folder `01-basic-fundamentals`
2. Baca README.md di setiap folder
3. Implement code di setiap file .go
4. Practice dengan membuat mini projects
5. Move to next folder setelah comfortable

## 📖 Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)

Happy coding! 🎉
EOF

# Buat script untuk generate boilerplate code
cat > generate-boilerplate.sh << 'EOF'
#!/bin/bash

# Script untuk generate boilerplate code template
echo "Generating boilerplate code templates..."

find . -name "*.go" -type f | while read file; do
    if [ ! -s "$file" ]; then  # If file is empty
        cat > "$file" << 'GOTEMPLATE'
package main

import (
	"fmt"
)

// TODO: Implement the functionality for this concept
// Refer to the README.md in this directory for what to learn

func main() {
	fmt.Println("Learning:", "REPLACE_WITH_CONCEPT_NAME")
	
	// TODO: Add your implementation here
}

// Example function - replace with actual implementation
func exampleFunction() {
	// Implementation goes here
}
GOTEMPLATE
    fi
done

echo "Boilerplate code generated!"
EOF

chmod +x generate-boilerplate.sh

echo "✅ Struktur pembelajaran Golang Standard Library berhasil dibuat!"
echo ""
echo "📁 Total struktur:"
echo "   - 11 kategori utama"
echo "   - 30+ packages/subpackages"  
echo "   - 100+ file latihan"
echo "   - README untuk setiap section"
echo ""
echo "🎯 Cara mulai belajar:"
echo "   1. cd go-stdlib-learning"
echo "   2. Mulai dari folder 01-basic-fundamentals"
echo "   3. Baca README.md di setiap folder"
echo "   4. ./generate-boilerplate.sh untuk template code"
echo ""
echo "💡 Tips: Ikuti urutan folder (01, 02, 03...) untuk pembelajaran yang optimal!"
echo ""
echo "Happy Learning! 🚀"