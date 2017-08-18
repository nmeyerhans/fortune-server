# Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# 	You may obtain a copy of the License at

# http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM debian:stretch

RUN apt-get update && apt-get -y install curl fortune-mod fortunes-bofh-excuses fortunes-off  && apt-get clean

RUN curl -sL -o /tmp/fortunes-spam.deb http://mirrors.cat.pdx.edu/ubuntu/pool/universe/f/fortunes-spam/fortunes-spam_1.8-0ubuntu1_all.deb && dpkg -i /tmp/fortunes-spam.deb

COPY fortune-server /

ENTRYPOINT "/fortune-server"
