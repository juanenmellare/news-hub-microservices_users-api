mockery --all --keeptree --case=snake --output .test/mocks
cp -R .test/mocks/internal/* test/mocks/
rm -rf .test/mocks/internal