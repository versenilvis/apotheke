<div align="center">
  <h1>Apotheke</h1>
  <p><b>Apotheke (/a.poˈteː.kə/) là công cụ alias lệnh thông minh.</b></p>
</div>

<div align="center">
  
  [![Stars](https://img.shields.io/badge/Stars-000?style=for-the-badge&logo=github&logoColor=white&labelColor=000000)](https://github.com/versenilvis/apotheke/stargazers)
  [![Twitter](https://img.shields.io/badge/Follow_me-000?style=for-the-badge&logo=x&logoColor=white&labelColor=000000)](https://x.com/VerseNilVis)

</div>

<div align="center">

  [![License: AGPL-3.0 license](https://img.shields.io/badge/License-AGPL_v3-blue?style=for-the-badge&logo=github&logoColor=white)](./LICENSE.md)
  [![Status](https://img.shields.io/badge/status-beta-yellow?style=for-the-badge&logo=github&logoColor=white)]()
  [![Documentation](https://img.shields.io/badge/docs-available-brightgreen?style=for-the-badge&logo=github&logoColor=white)](./docs/README.md)
  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen?style=for-the-badge&logo=github&logoColor=white)](./CONTRIBUTING.md)

</div>

Tiếng Việt | [English](./README.md)

> [!WARNING]
> **Hiện tại, Apotheke đang trong quá trình phát triển.**

## Xem trước

<img width="2531" height="742" alt="image" src="https://github.com/user-attachments/assets/cba2cc36-aa43-468d-994e-fc86cfb77c4f" />

## Cài đặt

**Một dòng lệnh (khuyên dùng):**

```bash
curl -sS https://raw.githubusercontent.com/versenilvis/apotheke/main/install.sh | sh
```

**Hoặc với Go:**

```bash
go install github.com/versenilvis/apotheke/cmd/apotheke@latest
```

**Hoặc build từ source:**

```bash
git clone https://github.com/versenilvis/apotheke
cd apotheke
make install
```

## Cấu hình Shell

Thêm vào file config shell để kích hoạt shortcut `a`:

Bash:

```bash
echo 'eval "$(apotheke init bash)"' >> ~/.bashrc
```

Zsh:

```bash
echo 'eval "$(apotheke init zsh)"' >> ~/.zshrc
```

Fish:

```fish
echo 'apotheke init fish | source' >> ~/.config/fish/config.fish
```

## Các lệnh

### add

Thêm bookmark cho lệnh mới.

```
a add <tên> <lệnh> [flags]
```

| Flag            | Mô tả                                                                                |
| --------------- | ------------------------------------------------------------------------------------ |
| `--cwd <path>`  | Đặt thư mục làm việc. Lệnh sẽ `cd` đến thư mục này trước khi chạy.                   |
| `--tags <tags>` | Các tag phân cách bằng dấu phẩy. Tag `danger` tự động bật xác nhận.                  |
| `--confirm`     | Luôn yêu cầu xác nhận trước khi chạy lệnh này.                                       |

<details>
  <summary>Ví dụ:</summary>

**Khuyên dùng** (đặt lệnh trong quotes):

```bash
a add ax "codex resume 019bc1e9-fb36-7f12-957f-061a532a9265"
a add kdp "kubectl delete pod"
a add deploy "make deploy" --confirm
a add build "npm run build" --cwd ~/project
a add prune "docker system prune -af" --tags docker,danger
```

**Cũng hoạt động** (không quotes, tất cả args sau tên trở thành lệnh):

```bash
a add ax codex resume 019bc1e9-fb36-7f12-957f-061a532a9265
a add kdp kubectl delete pod
```

**Không nên** (tên lệnh có khoảng trắng giữa các từ):

```bash
a add ax shell codex resume 019bc1e9-fb36-7f12-957f-061a532a9265
a add kdp del kubectl delete pod
```

**Bắt buộc** (quote lệnh khi có ký tự đặc biệt):

```bash
a add apo "curl -sS https://raw.githubusercontent.com/versenilvis/apotheke/main/install.sh | sh"
```
</details>

---

### rm

Xóa bookmark lệnh.

```
a rm <tên>
```
<details>
  <summary>Ví dụ:</summary>

```bash
a rm kdp
```

**Không hoạt động** (fuzzy match):

```bash
a rm k
```
</details>

> [!NOTE]
> Để an toàn, apotheke chỉ xóa bookmark nếu bạn gõ đúng tên.

---

### list

Liệt kê tất cả lệnh đã lưu.

```
a list [flags]
```

| Flag          | Mô tả                               |
| ------------- | ----------------------------------- |
| `--tag <tag>` | Lọc lệnh theo tag.                  |
| `-q <query>`  | Tìm kiếm lệnh theo tên hoặc nội dung. |

<details>
  <summary>Ví dụ:</summary>

```bash
a list                # hiện tất cả
a list --tag docker   # chỉ hiện lệnh có tag docker
a list -q kubectl     # tìm "kubectl"
```
</details>

> [!TIP]
> Gõ 'a' để liệt kê tất cả lệnh.

---

### show

Hiện chi tiết của một lệnh cụ thể.

```
a show <tên>
```

<details>
  <summary>Ví dụ:</summary>

```bash
a show kdp
# Output:
# kdp
#   Command:   kubectl delete pod
#   Tags:      k8s,danger
#   Confirm:   true
#   Frequency: 5
#   Last used: 2026-01-15 10:30:00
```
</details>

---

### run (mặc định)

Chạy lệnh đã lưu. Đây là hành động mặc định khi gõ `a <query>`.

```
a <query> [args...]
```

| Flag        | Mô tả                                                              |
| ----------- | ------------------------------------------------------------------ |
| `--dry-run` | Hiện lệnh sẽ chạy, nhưng không thực thi.                           |
| `-y`        | Bỏ qua prompt xác nhận (cho lệnh yêu cầu xác nhận).                |

Các args sau `<query>` được nối vào cuối lệnh đã lưu.

<details>
  <summary>Ví dụ:</summary>

```bash
a kdp my-pod              # chạy: kubectl delete pod my-pod
a kd my-pod               # fuzzy match "kdp" -> kubectl delete pod my-pod
a --dry-run kdp my-pod    # hiện lệnh mà không chạy
a -y kdp my-pod           # bỏ qua prompt xác nhận
```
</details>

---

### init

In script khởi tạo shell. Dùng với `eval` để kích hoạt function `a`.

```
apotheke init <shell>
```

| Shell  | Mô tả             |
| ------ | ----------------- |
| `bash` | Script Bash       |
| `zsh`  | Script Zsh        |
| `fish` | Script Fish       |

<details>
  <summary>Ví dụ:</summary>

```bash
eval "$(apotheke init zsh)"
```

</details>

---

## Matching

Khi chạy `a <query>`, resolver tìm match tốt nhất:

| Ưu tiên | Loại   | Mô tả                              |
| ------- | ------ | ---------------------------------- |
| 1       | Exact  | Query khớp chính xác tên lệnh      |
| 2       | Prefix | Query là prefix của tên lệnh       |
| 3       | Fuzzy  | Query fuzzy-match với tên lệnh     |

> [!IMPORTANT]
> Nếu nhiều lệnh khớp, một interactive picker sẽ hiện lên.
> 
> Ranking dùng **frecency** = tần suất × độ mới. Lệnh dùng thường xuyên và gần đây sẽ xếp cao hơn.

## An toàn

Lệnh có thể nguy hiểm. Các tính năng an toàn:

| Tính năng          | Mô tả                                                     |
| ------------------ | --------------------------------------------------------- |
| Flag `--confirm`   | Lệnh luôn hỏi "Execute? [y/N]" trước khi chạy             |
| Tag `danger`       | Lệnh có tag "danger" luôn yêu cầu xác nhận                |
| `--dry-run`        | Xem trước lệnh mà không chạy                              |
| Interactive picker | Nhiều lệnh khớp yêu cầu chọn rõ ràng                      |

## Dữ liệu

| Mục      | Đường dẫn                                |
| -------- | ---------------------------------------- |
| Database | `~/.local/share/apotheke/apotheke.db`    |
| Ghi đè   | Đặt biến môi trường `XDG_DATA_HOME`      |


## Gỡ cài đặt

```bash
rm ~/.local/bin/apotheke
rm -rf ~/.local/share/apotheke
```

Xóa dòng `eval` trong file config shell của bạn.

<details>
  <summary><h2>FAQ</h2></summary>
  
### Q: Tại sao bạn làm cái này?
A: Để lưu lệnh 'codex resume' và 'cursor-agent --resume=' mà tôi hay quên sau khi tắt terminal.

### Q: Apotheke khác gì với shell aliases?
A: Shell aliases cố định và cần chỉnh sửa file config. Apotheke cung cấp:
- Fuzzy matching (a kd → kubectl delete pod)
- Frecency ranking (lệnh dùng thường xuyên xếp cao hơn)
- Tags và tổ chức
- Xác nhận an toàn cho lệnh nguy hiểm
- Nối argument (a kdp my-pod → kubectl delete pod my-pod)

### Q: Khác gì với shell history?
A: History tìm tất cả lệnh. Apotheke chỉ lưu lệnh bạn bookmark với tên có ý nghĩa.

### Q: Có hoạt động trên Windows không?
A: Có, nhưng shell integration cần Git Bash, WSL, hoặc PowerShell với setup riêng.

### Q: Có thể sync giữa các máy không?
A: Chưa có sẵn. Có thể trong tương lai. Hoặc bạn có thể copy file database thủ công.

### Q: "Apotheke" nghĩa là gì?
A: Tôi hỏi ChatGPT "storage" trong tiếng Hy Lạp cổ là gì và nó nói "Apotheke".
</details>

## Giấy phép

[AGPL-3.0 license](./LICENSE)

## Đóng góp

Vui lòng làm theo [Contributing](.github/CONTRIBUTING.md) khi tạo pull request.
