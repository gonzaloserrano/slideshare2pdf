# slideshare2pdf

I built this because some slideshare presentations can't be downloaded as PDF,
like the ones from the [Personal Finance for Engineers](https://cs007.blog/)
Standford course.

I want to read them in my Kindle so this program downloads the slides, converts
them to landscape and puts them together in a PDF.

## install

You need the [golang toolchain](https://golang.org/) to install the project. Then:

    go get github.com/gonzaloserrano/slideshare2pdf

## usage

    slideshare2pdf URL1 URL2...

### example

    ❯ slideshare2pdf 'https://www.slideshare.net/adamnash/stanford-cs-00701-personal-finance-for-engineers-introduction?ref=https://cs007.blog/' 'https://www.slideshare.net/adamnash/stanford-cs-00702-personal-finance-for-engineers-predictably-irrational-81475880?ref=https://cs007.blog/'
    2018/01/02 11:27:10 Processing https://www.slideshare.net/adamnash/stanford-cs-00701-personal-finance-for-engineers-introduction?ref=https://cs007.blog/
    2018/01/02 11:27:11 Gathering 29 images from https://www.slideshare.net/adamnash/stanford-cs-00701-personal-finance-for-engineers-introduction?ref=https://cs007.blog/
    2018/01/02 11:27:15 Creating output PDF
    2018/01/02 11:27:18 Successfully created stanford-cs-00701-personal-finance-for-engineers-introduction.pdf
    2018/01/02 11:27:18 Processing https://www.slideshare.net/adamnash/stanford-cs-00702-personal-finance-for-engineers-predictably-irrational-81475880?ref=https://cs007.blog/
    2018/01/02 11:27:19 Gathering 22 images from https://www.slideshare.net/adamnash/stanford-cs-00702-personal-finance-for-engineers-predictably-irrational-81475880?ref=https://cs007.blog/
    2018/01/02 11:27:21 Creating output PDF
    2018/01/02 11:27:23 Successfully created stanford-cs-00702-personal-finance-for-engineers-predictably-irrational-81475880.pdf

    ❯ ls -l *.pdf
    -rw-r--r-- 1 gonzalo staff 16123448 ene  2 11:27 stanford-cs-00701-personal-finance-for-engineers-introduction.pdf
    -rw-r--r-- 1 gonzalo staff 12097833 ene  2 11:27 stanford-cs-00702-personal-finance-for-engineers-predictably-irrational-81475880.pdf
