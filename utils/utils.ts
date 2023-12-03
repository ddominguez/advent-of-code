function getRootPath() {
  const proc = Bun.spawnSync(["git", "rev-parse", "--show-toplevel"]);
  return proc.stdout.toString().trim();
}

export const rootPath = getRootPath();
