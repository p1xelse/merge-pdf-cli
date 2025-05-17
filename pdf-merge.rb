class PdfMerge < Formula
  desc "Merge pdf utility"
  homepage "https://github.com/p1xelse/merge-pdf-cli"
  url "https://github.com/p1xelse/merge-pdf-cli/releases/download/v1.0.0/pdf-merge"
  sha256 "3b23ecc8ef5919d681227b778f905599c3a7b1101c6b6965bdd4c58b5dd01330"

  def install
    bin.install "pdf-merge"
  end

  test do
    system "#{bin}/pdf-merge", "--version"
  end
end