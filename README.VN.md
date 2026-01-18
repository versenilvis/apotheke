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
curl -sS https://raw.githubusercontent.com/versenilvis/apotheke/main/install.sh | sudo sh
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
> [!IMPORTANT]
> - Vui lòng đọc từ [docs](./docs/commands.md)

## Hướng dẫn
> [!IMPORTANT]
> - Vui lòng đọc từ [docs](./docs/commands.md)

## Ví dụ khác
> [!IMPORTANT]
> - Vui lòng đọc từ [docs](./docs/commands.md)

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
