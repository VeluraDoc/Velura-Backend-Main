#!/usr/bin/env python3
import argparse, sys, os
from pdf2docx import Converter

def pdf_to_docx(pdf_path: str, docx_path: str):
    cv = Converter(pdf_path)
    try:
        cv.convert(docx_path)
    finally:
        cv.close()

def main():
    p = argparse.ArgumentParser()
    p.add_argument("-i", "--input", required=True, help="Path to input PDF")
    p.add_argument("-o", "--output", required=True, help="Path to output DOCX")
    args = p.parse_args()

    if not args.input.lower().endswith(".pdf"):
        print("input must be .pdf", file=sys.stderr)
        sys.exit(2)

    try:
        pdf_to_docx(args.input, args.output)
        print("OK")
    except Exception as e:
        print(f"convert error: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()
