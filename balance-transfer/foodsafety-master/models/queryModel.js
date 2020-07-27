var queryModel = {};

queryModel.getBlockByNumber = function(peer, blockNumber, username, org) {
    //1 节点地址
    var target = buildTarget(peer, org);
    //2 channel 
    var channel = helper.getChannelForOrg(org);
    //3  获取当前注册用户
    return helper.getRegisteredUsers(username, org).then((member) => {
        //4 查询区块信息
        return channel.queryBlock(parseInt(blockNumber), target);
    }, (err) => {
        logger.info('Failed to get submitter "' + username + '"');
        return 'Failed to get submitter "' + username + '". Error: ' + err.stack ?
            err.stack : err;
    }).then((response_payloads) => {
        //5 处理结果
        if (response_payloads) {
            //logger.debug(response_payloads);
            logger.debug(response_payloads);
            return response_payloads; //response_payloads.data.data[0].buffer;
        } else {
            logger.error('response_payloads is null');
            return 'response_payloads is null';
        }
    }, (err) => {
        logger.error('Failed to send query due to error: ' + err.stack ? err.stack :
            err);
        return 'Failed to send query due to error: ' + err.stack ? err.stack : err;
    }).catch((err) => {
        logger.error('Failed to query with error:' + err.stack ? err.stack : err);
        return 'Failed to query with error:' + err.stack ? err.stack : err;
    });
};

module.exports = queryModel;