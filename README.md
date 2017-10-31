# aws-scripts
When you change profile of your cached objects, the slab sizes in memcache don't change, if you have multiple region setup with cluster in each region.
Rebooting the clusters becomes a bit of a chore. This repo is to script work of rebooting each node and veryfing some metrics visually.
The end goal is to move node by node per each region and refer to a TSDB to confirm that changes have no negative effects

