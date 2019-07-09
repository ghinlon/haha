# Automatic pair insertion

# Links

* [Workflow question: automatic pair insertion : vim](https://www.reddit.com/r/vim/comments/3wggs8/workflow_question_automatic_pair_insertion/) 


# Overview

* [GitHub - jiangmiao/auto-pairs: Vim plugin, insert or delete brackets, parens, quotes in pair](https://github.com/jiangmiao/auto-pairs)  

  Don't like this one, the readme is awful.  

* [GitHub - cohama/lexima.vim: Auto close parentheses and repeat by dot dot dot...](https://github.com/cohama/lexima.vim)

  Not good. a little stupid. It always input the pairs.

* [GitHub - Raimondi/delimitMate: Vim plugin, provides insert mode auto-completion for quotes, parens, brackets, etc.](https://github.com/Raimondi/delimitMate)  

  I choose this one.

# Others

From [Workflow question: automatic pair insertion : vim](https://www.reddit.com/r/vim/comments/3wggs8/workflow_question_automatic_pair_insertion/cy04nk2?utm_source=share&utm_medium=web2x)

```zsh
 % cat ~/.config/nvim/UltiSnips/autopair.snippets
snippet ((
      (${1:inner_text})${0}
snippet )) Parens with leader
      ${1}(${2:inner_text})${0}
```

Doesn't how to make this work. 


