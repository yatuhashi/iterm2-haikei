bg() {
    $GOPATH/bin/change
    if [ -z "$BUFFER" ]; then
            osascript -e "tell application \"iTerm\"
                                        tell current window
                                            tell current session
                                                set background image to \"$HOME/Pictures/Fate/Haikei/result.jpg\"
                                            end tell
                                        end tell
                                    end tell"
            zle reset-prompt
    else
        zle accept-line
    fi
}
zle -N bg
bindkey '^m' bg
