Track where to copy assets from.

Keep the packages listed in tools.go and track their version, they are
where we copy the files from.

Keep this go.mod separate from the containing project, so we can track
exactly what version we copied the files from.

To copy newer versions of the assets (while merging local edits),
create a branch from where the assets were last copied, bump the
version in this module, and re-run ./copy-files here. Commit the
resulting changes and merge back to master. This lets git deal with
merging edits.
