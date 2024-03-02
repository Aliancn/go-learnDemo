FROM ubuntu:latest
LABEL authors="aliancn"

ENTRYPOINT ["top", "-b"]

# 使用MySQL官方镜像作为基础镜像
FROM mysql:latest

# 设置环境变量，定义MySQL root用户的密码
ENV MYSQL_ROOT_PASSWORD=yourpassword

# 暴露MySQL默认端口
EXPOSE 3306