# README

Magnolia is a web application by Rails.

## Ruby version
```
sudo apt-get install -y git build-essential make libssl-dev libreadline-dev

git clone https://github.com/rbenv/rbenv.git ~/.rbenv
git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
git clone https://github.com/rbenv/rbenv-vars.git ~/.rbenv/plugins/rbenv-vars

# Modify your ~/.zshrc file instead of ~/.bash_profile
echo 'export PATH=$HOME/.rbenv/bin:$PATH' >> ~/.bashrc
echo 'eval "$(rbenv init -)"' >> ~/.bashrc

# After re-login
rbenv install -l
CONFIGURE_OPTS="--disable-install-doc" rbenv install 2.2.5
rbenv local 2.2.5
gem install bundler
```

## System dependencies
```
sudo apt-get install -y cmake libicu-dev sdcv pkg-config libpq-dev

# Install nodejs
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get update
sudo apt-get install -y nodejs
```

## Configuration
```
sudo apt-get install -y pwgen
vi .rbenv-vars
vi config/database.yml
vi config/sidekiq.yml
vi config/deploy/production.rb
```

## Database creation
```
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

## Database initialization
```
bundle exec rake db:migrate
bundle exec rake db:seed
bundle exec rake spree_sample:load
```

## How to run the test suite
```
rake test
```

## Services (job queues, cache servers, search engines, etc.)
```
sudo apt-get install -y postgresql redis-server nginx

# Install jdk8
wget http://download.oracle.com/otn-pub/java/jdk/8u101-b13/jdk-8u101-linux-x64.tar.gz
sudo tar xf jdk-8u101-linux-x64.tar.gz -C /opt/jdk8
echo 'JAVA_HOME="/opt/jdk8"' >> /etc/environment

# Install elasticsearch
wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
echo "deb https://packages.elastic.co/elasticsearch/5.x/debian stable main" | sudo tee -a /etc/apt/sources.list.d/elasticsearch-5.x.list
sudo apt-get update
sudo apt-get install -y elasticsearch

# Install elasticsearch-analysis-ik
git clone https://github.com/medcl/elasticsearch-analysis-ik.git
cd elasticsearch-analysis-ik
mvn package
sudo unzip target/releases/elasticsearch-analysis-ik-5.0.0-alpha5.zip -d /usr/share/elasticsearch/plugins/ik
sudo service elasticsearch restart
```

## Deployment instructions
```
# deploy
bundle exec cap production deploy
# upload puma.conf
bundle exec cap production puma:config
# upload nginx config file
bundle exec cap production puma:nginx_config
# create sitemap.xml.gz
bundle exec cap production deploy:sitemap:create
```

## Issues

### Peer authentication failed for user

Need edit file "/etc/postgresql/9.5/main/pg_hba.conf" change line:
```
    local   all             all                                     peer
```

to:
```
    local   all             all                                     md5
```

