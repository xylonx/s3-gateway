# s3-gateway

this project is used to be compatible with any other storage type.

there are tons of object storage services in the internat. However, many of them are compatible with AWS S3 standard. For convenience, it's worthy to using a gateway that exposes the S3 standard API for consumer and saves the real data to the 'real' backend like [AWS](https://aws.amazon.com/s3/), [Aliyun OSS](https://cn.aliyun.com/product/oss), [Tencentcloud COS](https://intl.cloud.tencent.com/product/cos) and so on. 

Further, save file to those that is not compatible with the S3 standard. In this perspective, it is just like a individual object storage service like [MinIO](https://min.io/).