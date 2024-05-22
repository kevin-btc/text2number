from transforms import alpha2digit
import sys

def main():
    if len(sys.argv) != 2:
        print("Usage: python search_and_replace_by_num.py <sentence>")
        sys.exit(1)

    converted_sentence = alpha2digit(sys.argv[1], "fr", ordinal_threshold=0)
    print(converted_sentence)

if __name__ == "__main__":
    main()