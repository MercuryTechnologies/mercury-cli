{
  description = "Mercury CLI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils/v1.0.0";
  };

  nixConfig = {
    extra-substituters = [ "https://cache.mercury.com" ];
    extra-trusted-public-keys = [
      "cache.mercury.com:yhfFlgvqtv0cAxzflJ0aZW3mbulx4+5EOZm6k3oML+I="
    ];
  };

  outputs = { self, nixpkgs, flake-utils }:
    let
      mkMercury = pkgs:
        let
          version = self.shortRev or self.dirtyShortRev or "dev";
          commit = self.rev or self.dirtyRev or "unknown";
        in
        pkgs.buildGoModule {
          pname = "mercury";
          inherit version;
          src = self;
          # When go.sum changes: set to pkgs.lib.fakeHash, run `nix build`,
          # copy the sha256 from the error, paste it here.
          vendorHash = "sha256-aWtPggGsjcmLA7ESBbKVXegsVGflZjJ4BvNDIafg5+A=";
          subPackages = [ "cmd/mercury" ];
          env.CGO_ENABLED = "0";
          ldflags = [
            "-s"
            "-w"
            "-X main.version=${version}"
            "-X main.commit=${commit}"
          ];
          meta = with pkgs.lib; {
            description = "Mercury CLI";
            homepage = "https://github.com/MercuryTechnologies/mercury-cli";
            license = licenses.asl20;
            mainProgram = "mercury";
          };
        };
    in
    {
      overlays.default = final: _prev: {
        mercury-cli = mkMercury final;
      };
    } // flake-utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ] (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        mercury = mkMercury pkgs;
      in
      {
        packages = {
          default = mercury;
          mercury-cli = mercury;
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go_1_25
            gopls
            gotools
            goreleaser
          ];
        };
      });
}
