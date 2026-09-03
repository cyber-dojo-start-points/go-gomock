
lambda { |stdout,stderr,status|
  output = stdout + stderr

  # go test prints one status line per package. These are its own words for a
  # package that never reached its tests, such as one that does not compile.
  return :amber if /^FAIL\b.*\[build failed\]/.match(output)
  return :amber if /^FAIL\b.*\[setup failed\]/.match(output)

  # The go runtime writes these at the start of a line when a test blows up
  # rather than failing an assertion. go test reports such a package as FAIL,
  # so an exception has to be recognised before the red below.
  return :amber if /^panic: /.match(output)
  return :amber if /^fatal error: /.match(output)

  # A package whose tests ran and passed prints ok, one whose tests ran and
  # failed prints FAIL followed by the package name. A package holding no test
  # files prints ? instead, so a kata where nothing ran matches neither and
  # falls through to the amber below.
  return :green if status == 0 && /^ok\s/.match(output)
  return :red   if status != 0 && /^FAIL\s/.match(output)

  # mockgen refusing a source file lands here: go test never runs, so there is
  # no package status line to read at all.
  return :amber
}
