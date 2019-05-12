# cliassert

## Description

Simple command line assertion tool.

## Synopsis

```text
cliassert [ -assertion-case value ] command
```

## Examples

You can test the result of command execution.

```console
% cliassert -stdout-contain ok echo case-is-ok
% echo $?
0
```

If the test fails, details will be displayed on stderr.

```console
% cliassert -stdout-contain ok echo case-is-ng
[failure] stdout should contain ok.
% echo $?
1
```

Multiple tests can be set.
If the `-v` option is added, the details of the success case will also be displayed.

```console
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

You can pass the standard output of a command to standard output using the `pass` option.

```console
% cliassert -pass -exit-status 0 echo pass
pass
```

## Options

```text
-exit-status value
      String equal to exit-status
-pass
      Pass stdout of command to stdout
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
-v    Show verbose
```
