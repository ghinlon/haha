# vim

# help.txt

```
gO: to see the table of contets

ctrl-]/ctrl-t|ctrl-o 		jump to a subject 
:help word then hit ctrl-d = :helpgrep word
```

# what default key bind I never used

* Insert mode

* Normal mode

```
ctrl-h same as h
ctrl-j same as j
ctrl-n same as j
ctrl-p same as k
space same as l
```

# windows.txt

# tabpage.txt

```
:tabe : tabedit :tabnew 
	Open a new tab page with an empty window, after the current tab page.

:tabf :tabfind
	OPen a new tab page and edit {file} in 'path'

ctrl-w gf	Open a new tab page and edit the file name under the cursor.

:tabc :tabclose		" close the current tab page	
:tabo :tabonly		" close all tab pages except the current one
:tabn :tabnext gt	" Go to the next tab page.
:tabp :tabN gT		" Go to the previous tab page.
:tabr :tabrewind :tabfir :tabfirst	" Go to the first tab page.
:tabl :tablast				" Go to the last tab page.
:tabs		" List the tab pages and the windows they contain.
:tabm
:tabd :tabdo	" Execute {cmd} in each tab page
```

# tagsrch.txt

```
ctrl-t		Jump to older entry in the tag stack
:po :pop	Jump to older entry in tag stack
:ta :tag	Jump to newer entry in tag stack
:tags		Show the contents of th etag stack
:ts :tselect [name]	List the tags that match [name], using the information in the tags file.
sts :stselect		Does :tselect and splits the window for the selected tag.
```

# starting.txt

```
-R	Readonly mode.

2. Initialization

	The Nvim config file is named "init.vim", located at:
		Unix		~/.config/nvim/init.vim
		Windows		~/AppData/Local/nvim/init.vim
	Or if |$XDG_CONFIG_HOME| is defined:
				$XDG_CONFIG_HOME/nvim/init.vim

The $XDG_CONFIG_HOME and $XDG_DATA_HOME environment variables are used if they
exist, otherwise default values (listed below) are used.

CONFIG DIRECTORY (DEFAULT) ~
                  *$XDG_CONFIG_HOME*            Nvim: stdpath("config")
    Unix:         ~/.config                   ~/.config/nvim
    Windows:      ~/AppData/Local             ~/AppData/Local/nvim

DATA DIRECTORY (DEFAULT) ~
                  *$XDG_DATA_HOME*              Nvim: stdpath("data")
    Unix:         ~/.local/share              ~/.local/share/nvim
    Windows:      ~/AppData/Local             ~/AppData/Local/nvim-data

Note: Throughout the user manual these defaults are used as placeholders, e.g.
"~/.config" is understood to mean "$XDG_CONFIG_HOME or ~/.config".
```

# editing.txt

```
ctrl-g		Prints the current file name
g ctrl-g	Prints the current position of the cursor in file ways
:f {name}	Sets the current file name to {name}.
:0f		Remove the name of the current buffer.
:ls 		List all the currently known file names.
:e		Eedit the current file.
:e!		Edit the current file always.
:e {file} :e #[count]
:fin {file}	Find {file} in 'path' and then :edit it.

fileformat: unix, mac, dos

change file format:

:e file
:set fileformat=unix
:w

:n 		Edit next file.
:N :prev	Edit previous file
:rew :fir	Start editing the first file in the argument list.
:la :last	Start editing the last file in the argument list.
:wn	Write current file and start editing the next file.
:wN
:wp
:wa :wall	Write all changed buffers.
ZZ		Write current file, if modified and quit. ( same as ":x" ).
ZQ		Quit without checking for changes ( same as ":q!" ).

:cd		Change the current directory to the home directory
:pwd		Print the current directory on all systems.
:cd -		Change to the previous current directory.
```

# motion.txt

```
c	change
d	delete
y	yank
~ g~	swap case
gu	lowercase
gU	uppercase
gq	text formatting
g?	ROT13 encoding
>	shift right
<	shift left
zf	define a fold	多行会重叠起的

f/F{char}	To occurrence of {char} to the right/left.
t{char}	Till before [count]'th occurrence of {char} to the right.
T{char}	Till after [count]'th occurrence of {char} to the left.

; Repeat latest f, t, F or T
, Repeat latest f, t, F or T in opposite direction [count] times.

%	Go to {count percentage in the file.

(	[count] sentences backward.
)	[count] sentences forward.
{	[count] paragraphs backward.
}	[count] paragraphs forward.
]]	[count] sections forward or to the next '{' in the first column.
][	[count] sections forward or to the next '}' in the first column.
[[	[count] sections backward or to the previous '{' in the first column.
[]	[count] sections backward or to the previous '}' in the first column.

7. Marks

Jumping to a mark can be done in two ways:
1. With ` (backtick):	  The cursor is positioned at the specified location
			  and the motion is |exclusive|.
2. With ' (single quote): The cursor is positioned on the first non-blank
			  character in the line of the specified location and
			  the motion is linewise.

m{a-zA-Z}		Set mark at cursor position
m' or m`		Set the previous context mark.
m[ or m]		Set the '[ or ']mark. Useful when an operator is to be 
			simulated by multiple commands.
m< or m>		Set the '< or '> mark. Useful to change what the
			gv command selects.  做这2个标记后，gv就有法选中这块
:marks			List all the current marks.

8. Jumps

ctrl-o			Go to [count] Older cursor position in jump list
ctrl-i <Tab>		Go to [count] newer cursor position in jump list
:ju			Print the jump list
:cle			Clear the jump list of the current window.

g;			Go to [count] older position in change list
g,			Go to [count] newer cursor position in change list
:changes		Print the change list

9. Various motions

%			Find the next item in this line after or under the cursor and jump to its match.
```

# insert.txt

```
ctrl-a			Insert previously inserted text and stop insert.
ctrl-h <BS>		Delete the character before the cursor.
ctrl-w			Delete the word before the cursor
ctrl-u			Delete all entered characters before the cursor in the current line.
ctrl-i <Tab>		Insert a tab.
ctrl-j <NL>		Begin new line.
ctrl-m <CR>		Begin new line.
ctrl-n			Find next keyword
ctrp-p			Find previous keyword
ctrl-r			Intesert the contents of a register
ctrl-t			Insert one shiftwidth of indent at the start of the current line.
ctrl-d			Delete on shiftwidth of indent at the start of the current line.
0 ctrl-d		Delete all indent in the current line.
^ ctrl-d		Delete all indent in the current line. The indent is restored in the next line.
ctrl-v			Insert next non-digit literally.
ctrl-e			Insert the character which is below the cursor.
ctrl-y			Insert the character which is ablove the cursor.
ctrl-]			Trigger abbreviation.

7. Insert mode completion

ctr-x ctrl-l		Whole lines
ctr-x ctrl-n		Keywords in the current file
...
```

