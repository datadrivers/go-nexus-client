#!/bin/bash
SED="sed"

wget http://127.0.0.1:8081/service/rest/swagger.json -O ./nexus-swagger.json
${SED} -ri 's/createRepository"/createGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/getRepository"/getMavenGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/updateRepository"/updateMavenGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_1"/MavenHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_2"/MavenProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_3"/RawGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_4"/RawHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_5"/RawProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_6"/NpmGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_7"/NpmHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_8"/NpmProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_9"/HelmHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_10"/HelmProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_11"/PypiGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_12"/PypiHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_13"/PypiProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_14"/NugetGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_15"/NugetHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_16"/NugetProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_17"/RubygemsGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_18"/RubygemsHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_19"/RubygemsProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_20"/DockerGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_21"/DockerHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_22"/DockerProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_23"/YumGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_24"/YumHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_25"/YumProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_26"/AptHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_27"/AptProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_28"/CocoapodsProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_29"/ConanProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_30"/CondaProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_31"/GoGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_32"/GoProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_33"/P2ProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_34"/RGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_35"/RHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_36"/RProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_37"/BowerGroupRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_38"/BowerHostedRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_39"/BowerProxyRepository"/g' ./nexus-swagger.json
${SED} -ri 's/Repository_40"/GitlfsHostedRepository"/g' ./nexus-swagger.json

swagger-codegen generate \
    -i ./nexus-swagger.json \
    -l go \
    -o . \
    -c ./codegen-config.json \
    --git-user-id datadrivers \
    --git-repo-id go-nexus-client