# golang_skeleton

## 1. Thiết lập môi trường phát triển trên Windows

| WSL2 | Golang | VSCode + Extension | Docker | Make | TablePlus |
| - | - | - | - | - | - |

#### 1.1 Cài đặt WSL2
- Chuột phải vào thanh Start -> Mở Powershell (Admin): </br>
<code>wsl --install</code> </br>

- Khởi động lại

#### 1.2 Cài đặt Golang
- Link tải Golang: <link href="https://go.dev/doc/install">https://go.dev/doc/install</link> </br>
- Nhấn chuột phải vào file cài đặt -> Run as administrator (chạy với quyền quản trị) </br>

#### 1.3 Cài đặt VSCode + Extension
- Link tải VSCode: <link href="https://code.visualstudio.com/download">https://code.visualstudio.com/download</link></br>
- Nhấn chuột phải vào file cài đặt -> Run as administrator (chạy với quyền quản trị) </br>
- Cài đặt Extension: </br>
    <ul>
      <li>Auto Rename Tag</li>
      <li>Go</li>
      <li>Makefile Tools</li>
      <li>Material Icon Theme</li>
      <li>Docker</li>
      <li>WSL</li>
    </ul>
    
#### 1.4 Cài đặt Docker
- Link tải Docker: <link href="https://docs.docker.com/desktop/install/windows-install/">https://docs.docker.com/desktop/install/windows-install/</link></br>
- Chạy quyền quản trị file
- Khởi động lại

#### 1.5 Cài đặt Make
- Link tải Make: <link href="https://sourceforge.net/projects/mingw/files/Installer/mingw-get-setup.exe/download">https://sourceforge.net/projects/mingw/files/Installer/mingw-get-setup.exe/download</link></br>

#### 1.6 Cài đặt TablePlus
- Link tải TablePlus: <link href="https://tableplus.com/download">https://tableplus.com/download</link></br>
- Chọn phiên bản HĐH


## 2. Bắt đầu cơ bản với cấu trúc thư mục Cobra go
- Link tài liệu : <link href="https://github.com/spf13/cobra/blob/main/site/content/user_guide.md">https://github.com/spf13/cobra/blob/main/site/content/user_guide.md</link> </br>

- Ứng dụng dựa trên Cobra sẽ tuân theo cơ cấu tổ chức sau: </br>
  ▾ appName/
    ▾ cmd/
        add.go
        your.go
        commands.go
        here.go
      main.go

#### 2.1 Bắt đầu tạo cấu trúc thư mục và main.go

- Tạo thư mục theo tổ chức sau:
  ▾ apps/
    ▾ cmd/
        server/
          server.go
        root.go
      main.go
- Khai báo package cho toàn bộ file go
- Mở CMD hoặc Powershell lên: </br>
  cd tới thư mục dự án đã clone từ github: </br>
  - Đi tới thư mục apps: <code>cd apps</code> </br>
  - Khởi tạo dự án Go (gocms): <code>go mod init gocms</code> có thể thay gocms bằng tên dự án của mình, <code>go mod tidy</code> lệnh tải các package dự án