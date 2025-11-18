# Forensic filesystem

As a forensic tool, all write access to the examined files is **prohibited** by different technical measures. All files will be lazy loaded readonly into the main memory by mapping them upon their first view.

> To also prevent the writing of output to the current filesystem, use the `--readonly` flag or mount the filesystem as readonly. The software will still stay functional with these limitations.

## Multicore operations
All processor heavy operations, like searching or formating, will be done via multicore data handling and cached for faster response times. These operations are optimized for files that contain **one million** or more lines.
