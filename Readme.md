# Basic WC

The basic word counter is a straightforward command line utility designed to tally words provided through the standard input (STDIN) connection. Its primary function is to count the number of words, but it also offers additional capabilities based on specific flags. By default, it counts words, but if the -l flag is included, it switches to counting lines. Similarly, if the -b flag is provided, it changes the count to bytes instead.

---

## Usage

The `basic-wc` command line utility can be used as described below.

#### Default: Word Counting

The default functionality of basic-wc is to count the words in an input passed through STDIN. Here is an example of how to use it:

```bash
echo "This is a test sentence" | ./basic-wc
```

The output of this command will be `5`, as there are five words in the sentence.

#### Line Counting

If you wish to count the number of lines instead of words, use the `-l` flag. Here's how to use it:

```bash
echo -e "This is a test sentence.\nThis is another test sentence." | basic-wc -l
```

The output for this command will be `2`, as there are two lines in the input.

#### Byte Counting

To count the number of bytes in your input, you can use the -b flag. Here's an example of how to use it:

```bash
echo "This is a test sentence" | basic-wc -b
```

The output of this command will be the number of bytes in the provided sentence.

Keep in mind, your input can be a text file or any other source that can be piped into the command, not just text echoed directly in the command line.
