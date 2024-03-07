# Terminal beeping

The objective of this program is to make the terminal beep.

When a numeric key is pressed the terminal should beep as many times as the numeric input, else it should not do anything. The program should exit with a SIGINT (^C) and behaving as before the app ran.

This is meant to:

- Show that even ASCII is a binary encoding
- Expose concepts related to the tty interface

ASCII is a binary encoding of both familiar characters (e.g. `A`, `5`) and characters like `null`, `backspace`
