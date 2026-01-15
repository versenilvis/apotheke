package shell

import "fmt"

func InitBash() string {
	return `
function a() {
    if [[ $# -eq 0 ]]; then
        command apotheke list
        return
    fi
    
    case "$1" in
        add|rm|remove|list|ls|edit|show|init|help|--help|-h|version|--version|-v)
            command apotheke "$@"
            ;;
        *)
            command apotheke "$@"
            ;;
    esac
}
`
}

func InitZsh() string {
	return `
function a() {
    if [[ $# -eq 0 ]]; then
        command apotheke list
        return
    fi
    
    case "$1" in
        add|rm|remove|list|ls|edit|show|init|help|--help|-h|version|--version|-v)
            command apotheke "$@"
            ;;
        *)
            command apotheke "$@"
            ;;
    esac
}
`
}

func InitFish() string {
	return `
function a
    if test (count $argv) -eq 0
        command apotheke list
        return
    end
    
    switch $argv[1]
        case add rm remove list ls edit show init help --help -h version --version -v
            command apotheke $argv
        case '*'
            command apotheke $argv
    end
end
`
}

func Init(shell string) (string, error) {
	switch shell {
	case "bash":
		return InitBash(), nil
	case "zsh":
		return InitZsh(), nil
	case "fish":
		return InitFish(), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s (supported: bash, zsh, fish)", shell)
	}
}
