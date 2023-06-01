#!/bin/bash
cd /Users/wiloon/workspace/projects/newbee/newbee-mall-api-go || exit
package_name="mall-api"
echo "building"
rm -f mall-api
GOOS=linux GOARCH=amd64 go build -o ${package_name} main.go
cp ${package_name} /Users/wiloon/tmp

# mall-admin
cd /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin || exit

rm -rf dist

sed -i '' 's/mall-admin.wiloon.com/dbyhh-admin.ehoneycomb.net/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/index.html
sed -i '' 's/mall-admin.wiloon.com/dbyhh-admin.ehoneycomb.net/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/config/index.js
sed -i '' 's/mall-admin.wiloon.com/dbyhh-admin.ehoneycomb.net/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/src/utils/index.js
sed -i '' 's/mall.wiloon.com/dbyhh.ehoneycomb.net/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/src/views/ShopOrder.vue

npm run build:release
package_name="newbee-mall-admin.tar.gz"
gtar zcvf ${package_name} -C dist .
cp ${package_name} /Users/wiloon/tmp

sed -i '' 's/dbyhh-admin.ehoneycomb.net/mall-admin.wiloon.com/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/index.html
sed -i '' 's/dbyhh-admin.ehoneycomb.net/mall-admin.wiloon.com/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/config/index.js
sed -i '' 's/dbyhh-admin.ehoneycomb.net/mall-admin.wiloon.com/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/src/utils/index.js
sed -i '' 's/dbyhh.ehoneycomb.net/mall.wiloon.com/' /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-admin/src/views/ShopOrder.vue


# app
cd /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-app || exit
rm -rf dist

sed -i '' 's/mall.wiloon.com/dbyhh.ehoneycomb.net/'  /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-app/src/utils/axios.js

npm run build
package_name="newbee-mall-app.tar.gz"
gtar zcvf ${package_name} -C dist .
cp ${package_name} /Users/wiloon/tmp
sed -i '' 's/dbyhh.ehoneycomb.net/mall.wiloon.com/'  /Users/wiloon/workspace/projects/newbee/newbee-mall-vue3-app/src/utils/axios.js

