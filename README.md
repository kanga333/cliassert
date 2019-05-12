# cliassert

## Description

Simple command line assertion tool.

## Synopsis

```text
cliassert [ -assertion-case value ] command
```

```console
# You can check the result of command execution.
% cliassert -stdout-contain ok echo case-is-ok
% echo $?
0

# If the test fails, details will be displayed on stderr.
% cliassert -stdout-contain ok echo case-is-ng
[failure] stdout should contain ok. ()
% echo $?
1

# Multiple tests can be set.
# If the -v option is added, the details of the success case will also be displayed.
% cliassert -v -stdout-contain case -stdout-contain ok echo case-is-ng
[exit-status] 0

[stdout]
case-is-ng

[stderr]

---
[success] stdout should contain case.
[failure] stdout should contain ok.

2 cases, 1 failures.
```

## Options

```text
-exit-status value
      String equal to exit-status
-stderr-contain value
      String contained in stderr
-stderr-match value
      Regex matching stderr
-stderr-not-contain value
      String not contained in stderr
-stderr-not-match value
      Regex not matching stderr
-stdout-contain value
      String contained in stdout
-stdout-match value
      Regex matching stdout
-stdout-not-contain value
      String not contained in stdout
-stdout-not-match value
      Regex not matching stdout
-v    show verbose
```
