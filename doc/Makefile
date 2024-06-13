PDF = coursework.pdf
MAIN = coursework
FLAGS = -shell-escape
CC = pdflatex
SHIT = *.aux *.log *.out *.toc *.bbl *.blg *.xml *.bcf

all: bib

$(PDF) : $(MAIN).tex *.tex
	make resurrect
	-$(CC) $(FLAGS) $<
	-$(CC) $(FLAGS) $<
	-mv $(MAIN).pdf $(PDF)
	make moveout

# bib: $(MAIN).tex *.tex
# 	make resurrect
# 	-$(CC) $(FLAGS) $<  > /dev/null
# 	-bibtex $(MAIN).aux > /dev/null
# 	-$(CC) $(FLAGS) $<  > /dev/null
# 	-$(CC) $(FLAGS) $<  > /dev/null
# 	-mv $(MAIN).pdf $(PDF)
# 	make moveout

bib: $(MAIN).tex *.tex
	make resurrect
	-$(CC) $(FLAGS) $<  > /dev/null
	-$(CC) $(FLAGS) $<  > /dev/null
	-biber $(MAIN).bcf
	-$(CC) $(FLAGS) $<  > /dev/null
	-mv $(MAIN).pdf $(PDF)
	make moveout

moveout:
	-mkdir -p out > /dev/null
	-mv $(SHIT) out > /dev/null

resurrect:
	-mv out/* . > /dev/null

clean:
	-rm  $(SHIT) > /dev/null
	rm -rf out

.PHONY: all, clean, bib, moveout, resurrect
