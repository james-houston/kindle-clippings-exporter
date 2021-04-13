# Kindle clippings local exporter
Help automate the process of removing clippings from a kindle, and then send them in an email for review later.

Run this while a kindle is connected to your computer (tested with connection to a Raspberry Pi running this in cron). This program will handle discovering the connected kindle, reading the clippings file, parsing the clippings, finding which ones are new, and then emailing those clippings to you.

The clippings are parsed and stored on disk in the format `./KINDLE_ID/Book_Title/Clipping_Timestamp.txt`. These .txt files are how the program knows which clippings have been discovered before.

## Motivation
I leave my kindle disconnected from wifi, so when I want to get my clippings off the device I normally have to manually copy of _My Clippings.txt_ and read through the file. I wanted a simpler way to receive them and be able to review my most recent additions.

## Hardware setup
This has only been tested on my Kindle Paperwhite 3 from 2015. No idea how other versions of the kindle will work with this. I use a raspberry pi connected to the kindle via micro usb.

## Use
Clone this repo and copy the `EXAMPLE.env` file to `.env` and update with your values. If you do not want to send an email and only want .txt files on disk, skip this step. Right now only sending from gmail emails is supported. If you want to send from a different provider you will have to edit the smtp configuration in kindle_emailer.go

From there just connect your kindle and run `go run main.go`. The clippings should be discovered and written to disk.

## DISCLAIMER
Use at your own risk. I am not responsible for whatever happens to you. This program should help significantly reduce your risk of reading-related papercuts, but dramatically increases your chances of all other ailments.