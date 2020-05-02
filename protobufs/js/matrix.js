/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.Matrix');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');

goog.forwardDeclare('proto.MatrixType');

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
proto.Matrix = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Matrix, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.Matrix.displayName = 'proto.Matrix';
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
proto.Matrix.prototype.toObject = function(opt_includeInstance) {
  return proto.Matrix.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Matrix} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Matrix.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    device: jspb.Message.getFieldWithDefault(msg, 2, 0),
    channel: jspb.Message.getFieldWithDefault(msg, 3, 0),
    width: jspb.Message.getFieldWithDefault(msg, 4, 0),
    height: jspb.Message.getFieldWithDefault(msg, 5, 0),
    type: jspb.Message.getFieldWithDefault(msg, 6, 0),
    coloring: jspb.Message.getFieldWithDefault(msg, 7, ""),
    offset: jspb.Message.getFieldWithDefault(msg, 8, 0),
    brightness: jspb.Message.getFieldWithDefault(msg, 9, 0)
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
 * @return {!proto.Matrix}
 */
proto.Matrix.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Matrix;
  return proto.Matrix.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Matrix} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Matrix}
 */
proto.Matrix.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {number} */ (reader.readUint64());
      msg.setDevice(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setChannel(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setWidth(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setHeight(value);
      break;
    case 6:
      var value = /** @type {!proto.MatrixType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setColoring(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setOffset(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setBrightness(value);
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
proto.Matrix.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Matrix.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Matrix} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Matrix.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getDevice();
  if (f !== 0) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = message.getChannel();
  if (f !== 0) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = message.getWidth();
  if (f !== 0) {
    writer.writeUint32(
      4,
      f
    );
  }
  f = message.getHeight();
  if (f !== 0) {
    writer.writeUint32(
      5,
      f
    );
  }
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getColoring();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getOffset();
  if (f !== 0) {
    writer.writeUint32(
      8,
      f
    );
  }
  f = message.getBrightness();
  if (f !== 0) {
    writer.writeUint32(
      9,
      f
    );
  }
};


/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.Matrix.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional uint64 device = 2;
 * @return {number}
 */
proto.Matrix.prototype.getDevice = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setDevice = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional uint64 channel = 3;
 * @return {number}
 */
proto.Matrix.prototype.getChannel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setChannel = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional uint32 width = 4;
 * @return {number}
 */
proto.Matrix.prototype.getWidth = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setWidth = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional uint32 height = 5;
 * @return {number}
 */
proto.Matrix.prototype.getHeight = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setHeight = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional MatrixType type = 6;
 * @return {!proto.MatrixType}
 */
proto.Matrix.prototype.getType = function() {
  return /** @type {!proto.MatrixType} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {!proto.MatrixType} value */
proto.Matrix.prototype.setType = function(value) {
  jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * optional string coloring = 7;
 * @return {string}
 */
proto.Matrix.prototype.getColoring = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.Matrix.prototype.setColoring = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional uint32 offset = 8;
 * @return {number}
 */
proto.Matrix.prototype.getOffset = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setOffset = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional uint32 brightness = 9;
 * @return {number}
 */
proto.Matrix.prototype.getBrightness = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.Matrix.prototype.setBrightness = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


