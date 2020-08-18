#!/bin/csh
set echo
set name = "main"



set this_latex = pdflatex
if( -e /usr/bin/acroread ) then
 set run = "acroread"
else
 set run = "open"
endif    
rm -f *.log *.out *.aux *.dvi *~ *.bbl *.blg
rm -f *.ps $name.pdf *.bak *.toc *.lof *.lot
$this_latex $name
if( $status != 0 ) then
    echo "done borked sompin. Exiting."
    exit
endif
bibtex $name
$this_latex $name
$this_latex $name
$this_latex $name

# the end
