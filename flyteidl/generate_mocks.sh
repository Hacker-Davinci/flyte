#!/bin/bash
set -e
set -x

mockery --dir=gen/pb-go/flyteidl/service/ --all --output=clients/go/admin/mocks --with-expecter
mockery --dir=gen/pb-go/flyteidl/datacatalog/ --name=DataCatalogClient --output=clients/go/datacatalog/mocks --with-expecter
mockery --dir=gen/pb-go/flyteidl/cacheservice/ --name=CacheServiceClient --output=clients/go/cacheservice/mocks --with-expecter
