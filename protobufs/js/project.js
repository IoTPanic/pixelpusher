/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.Project');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Project = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Project, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.Project.displayName = 'proto.Project';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Project.prototype.toObject = function(opt_includeInstance) {
  return proto.Project.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Project} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Project.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    created: jspb.Message.getFieldWithDefault(msg, 3, 0),
    lastUpdate: jspb.Message.getFieldWithDefault(msg, 4, 0),
    client: jspb.Message.getFieldWithDefault(msg, 5, ""),
    active: jspb.Message.getFieldWithDefault(msg, 6, false),
    frontendState: msg.getFrontendState_asB64()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Project}
 */
proto.Project.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Project;
  return proto.Project.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Project} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Project}
 */
proto.Project.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setCreated(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLastUpdate(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setClient(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setActive(value);
      break;
    case 7:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setFrontendState(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Project.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Project.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Project} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Project.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCreated();
  if (f !== 0) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = message.getLastUpdate();
  if (f !== 0) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = message.getClient();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getActive();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
  f = message.getFrontendState_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      7,
      f
    );
  }
};


/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.Project.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.Project.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.Project.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.Project.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional uint64 created = 3;
 * @return {number}
 */
proto.Project.prototype.getCreated = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.Project.prototype.setCreated = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional uint64 last_update = 4;
 * @return {number}
 */
proto.Project.prototype.getLastUpdate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.Project.prototype.setLastUpdate = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional string client = 5;
 * @return {string}
 */
proto.Project.prototype.getClient = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.Project.prototype.setClient = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional bool active = 6;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.Project.prototype.getActive = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 6, false));
};


/** @param {boolean} value */
proto.Project.prototype.setActive = function(value) {
  jspb.Message.setProto3BooleanField(this, 6, value);
};


/**
 * optional bytes frontend_state = 7;
 * @return {string}
 */
proto.Project.prototype.getFrontendState = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * optional bytes frontend_state = 7;
 * This is a type-conversion wrapper around `getFrontendState()`
 * @return {string}
 */
proto.Project.prototype.getFrontendState_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getFrontendState()));
};


/**
 * optional bytes frontend_state = 7;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getFrontendState()`
 * @return {!Uint8Array}
 */
proto.Project.prototype.getFrontendState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getFrontendState()));
};


/** @param {!(string|Uint8Array)} value */
proto.Project.prototype.setFrontendState = function(value) {
  jspb.Message.setProto3BytesField(this, 7, value);
};

