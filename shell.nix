let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-unstable";
  pkgs = import nixpkgs {config = {}; overlays = [];};
in

pkgs.mkShell {
  packages = with pkgs; [
    go
    git
    entr
    gocyclo
    codespell
    ineffassign
  ];
}
