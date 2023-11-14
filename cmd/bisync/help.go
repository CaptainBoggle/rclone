package bisync

import (
	"strconv"
	"strings"
)

func makeHelp(help string) string {
	replacer := strings.NewReplacer(
		"|", "`",
		"{MAXDELETE}", strconv.Itoa(DefaultMaxDelete),
		"{CHECKFILE}", DefaultCheckFilename,
		// "{WORKDIR}", DefaultWorkdir,
	)
	return replacer.Replace(help)
}

var shortHelp = `Perform bidirectional synchronization between two paths.`

var rcHelp = makeHelp(`This takes the following parameters

- path1 - a remote directory string e.g. |drive:path1|
- path2 - a remote directory string e.g. |drive:path2|
- dryRun - dry-run mode
- resync - performs the resync run
- checkAccess - abort if {CHECKFILE} files are not found on both filesystems
- checkFilename - file name for checkAccess (default: {CHECKFILE})
- maxDelete - abort sync if percentage of deleted files is above
  this threshold (default: {MAXDELETE})
- force - Bypass maxDelete safety check and run the sync
- checkSync - |true| by default, |false| disables comparison of final listings,
              |only| will skip sync, only compare listings from the last run
- createEmptySrcDirs - Sync creation and deletion of empty directories. 
			  (Not compatible with --remove-empty-dirs)
- removeEmptyDirs - remove empty directories at the final cleanup step
- filtersFile - read filtering patterns from a file
- ignoreListingChecksum - Do not use checksums for listings
- resilient - Allow future runs to retry after certain less-serious errors, instead of requiring resync. 
            Use at your own risk!
- workdir - server directory for history files (default: |~/.cache/rclone/bisync|)
- backupdir1 - --backup-dir for Path1. Must be a non-overlapping path on the same remote.
- backupdir2 - --backup-dir for Path2. Must be a non-overlapping path on the same remote.
- noCleanup - retain working files
- gui - Whether to enable the bisync GUI. (default: false)
- guiDir - Use custom dir to store GUI file. Accepts any rclone remote dir path (but local is recommended.) (default: --workdir)
- guiRefresh - Interval for GUI HTML page auto-refresh. (default: 30s)
- guiStats - Interval between printing stats to GUI, e.g. 500ms, 60s, 5m (0 to disable) (default: 30s)
- guiMaxRows - Max number of rows to keep and display on GUI. (default: 50)
- guiErrPop - Automatically open the GUI to alert if bisync errors. --gui-dir must be a local path. (default: false)

See [bisync command help](https://rclone.org/commands/rclone_bisync/)
and [full bisync description](https://rclone.org/bisync/)
for more information.`)

var longHelp = shortHelp + makeHelp(`

[Bisync](https://rclone.org/bisync/) provides a
bidirectional cloud sync solution in rclone.
It retains the Path1 and Path2 filesystem listings from the prior run.
On each successive run it will:
- list files on Path1 and Path2, and check for changes on each side.
  Changes include |New|, |Newer|, |Older|, and |Deleted| files.
- Propagate changes on Path1 to Path2, and vice-versa.

See [full bisync description](https://rclone.org/bisync/) for details.
`)
