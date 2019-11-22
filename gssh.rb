require "FileUtils"

class Gssh < Formula
  homepage "https://github.com/ssuareza/gssh"

  $v = "v0.0.1"
  url "https://github.com/ssuareza/gssh/releases/download/#{$v}/gssh-#{$v}-darwin-amd64"
  version $v
  sha256 "16252aaa702a9796cd552ed559ba5e3a93c4371aa35834b7c2771a63e58ce250"

  def install
    FileUtils.mv("gssh-#{$v}-darwin-amd64", "gssh")
    bin.install "gssh"
  end

  #test do
  #  system "#{bin}/gssh", "--version"
  #end
end