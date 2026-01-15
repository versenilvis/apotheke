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
            local __apotheke_result
            __apotheke_result="$(command apotheke resolve "$@")"
            local __apotheke_exit=$?
            
            if [[ $__apotheke_exit -eq 0 && -n "$__apotheke_result" ]]; then
                eval "$__apotheke_result"
            fi
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
            local __apotheke_result
            __apotheke_result="$(command apotheke resolve "$@")"
            local __apotheke_exit=$?
            
            if [[ $__apotheke_exit -eq 0 && -n "$__apotheke_result" ]]; then
                eval "$__apotheke_result"
            fi
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
            set -l __apotheke_result (command apotheke resolve $argv)
            set -l __apotheke_exit $status
            
            if test $__apotheke_exit -eq 0 -a -n "$__apotheke_result"
                eval $__apotheke_result
            end
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
