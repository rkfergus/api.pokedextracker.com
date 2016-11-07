'use strict';

const CreateBoomError = require('create-boom-error').bind(exports);

CreateBoomError('ExistingDex', 422, 'a dex with that URL already exists');
CreateBoomError('ExistingUsername', 422, 'username is already taken');
CreateBoomError('ForbiddenAction', 403, (action) => `${action} is not allowed`);
CreateBoomError('InvalidPassword', 422, 'password is invalid');
CreateBoomError('NotFound', 404, (item) => `${item} not found`);

exports.DuplicateKey = (err) => err.message.includes('duplicate key value');
