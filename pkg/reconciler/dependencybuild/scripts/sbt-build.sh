#!/usr/bin/env bash
set -eu
set -o pipefail
#fix this when we no longer need to run as root
export HOME=/root

TOOL_VERSION="$(params.TOOL_VERSION)"
export SBT_DIST="/opt/sbt/${TOOL_VERSION}"
echo "SBT_DIST=${SBT_DIST}"

if [ ! -d "${SBT_DIST}" ]; then
    echo "SBT home directory not found at ${SBT_DIST}" >&2
    exit 1
fi

export PATH="${SBT_DIST}/bin:${PATH}"

mkdir $(workspaces.source.path)/logs
mkdir $(workspaces.source.path)/packages
{{INSTALL_PACKAGE_SCRIPT}}

mkdir "$HOME/.sbt"
cat > "$HOME/.sbt/repositories" <<EOF
[repositories]
  local
  my-maven-proxy-releases: $(params.CACHE_URL), allowInsecureProtocol
EOF

mkdir "$HOME/.sbt/1.0/"
cat >"$HOME/.sbt/1.0/global.sbt" <<EOF
publishTo := Some(("MavenRepo" at s"file://$(workspaces.source.path)/hacbs-jvm-deployment-repo").withAllowInsecureProtocol(true)),
EOF
#This is replaced when the task is created by the golang code
cat <<EOF
Pre build script: {{PRE_BUILD_SCRIPT}}
EOF
{{PRE_BUILD_SCRIPT}}

cp -r $(workspaces.source.path)/workspace $(workspaces.source.path)/source

echo "Command is:"
echo "sbt $@ "

eval "sbt $@" | tee $(workspaces.source.path)/logs/sbt.log


tar -czf "$(workspaces.source.path)/hacbs-jvm-deployment-repo.tar.gz" -C "$(workspaces.source.path)/hacbs-jvm-deployment-repo" .
