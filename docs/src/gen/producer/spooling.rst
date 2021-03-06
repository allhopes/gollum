.. Autogenerated by Gollum RST generator (docs/generator/*.go)

Spooling
========

This producer is meant to be used as a fallback if another producer fails to
send messages, e.g. because a service is down. It does not really produce
messages to some other service, it buffers them on disk for a certain time
and inserts them back to the system after this period.




Parameters
----------

**Enable** (default: true)

  Switches this plugin on or off.
  

**Path** (default: /var/run/gollum/spooling)

  Sets the output directory for spooling files. Spooling files will
  be stored as "<path>/<stream name>/<number>.spl".
  By default this parameter is set to "/var/run/gollum/spooling".
  
  

**MaxFileSizeMB** (default: 512, unit: mb)

  Sets the size limit in MB that causes a spool file rotation.
  Reading messages back into the system will start only after a file is
  rotated.
  By default this parameter is set to 512.
  
  

**MaxFileAgeMin** (default: 1, unit: min)

  Defines the duration in minutes after which a spool file
  rotation is triggered (regardless of MaxFileSizeMB). Reading messages back
  into the system will start only after a file is rotated.
  By default this parameter is set to 1.
  
  

**MaxMessagesSec**

  Sets the maximum number of messages that will be respooled
  per second. Setting this value to 0 will cause respooling to send as fast as
  possible.
  By default this parameter is set to 100.
  
  

**RespoolDelaySec** (default: 10, unit: sec)

  Defines the number of seconds to wait before trying to
  load existing spool files from disk after a restart. This setting can be used
  to define a safe timeframe for gollum to set up all required connections and
  resources before putting additionl load on it.
  By default this parameter is set to 10.
  
  

**RevertStreamOnFallback** (default: false)

  This allows the spooling fallback to handle the
  messages that would have been sent back by the spooler if it would have
  handled the message. When set to true it will revert the stream of the
  message to the previous stream ID before sending it to the Fallback stream.
  By default this parameter is set to false.
  
  

**BufferSizeByte** (default: 8192)

  Defines the initial size of the buffer that is used to read
  messages from a spool file. If a message is larger than this size, the buffer
  will be resized.
  By default this parameter is set to 8192.
  
  

**Batch/MaxCount** (default: 100)

  defines the maximum number of messages stored in memory before
  a write to file is triggered.
  By default this parameter is set to 100.
  
  

**Batch/TimeoutSec** (default: 5, unit: sec)

  defines the maximum number of seconds to wait after the last
  message arrived before a batch is flushed automatically.
  By default this parameter is set to 5.
  
  

Parameters (from components.RotateConfig)
-----------------------------------------

**Rotation/Enable** (default: false)

  If this value is set to "true" the logs will rotate after reaching certain thresholds.
  By default this parameter is set to "false".
  
  

**Rotation/TimeoutMin** (default: 1440, unit: min)

  This value defines a timeout in minutes that will cause the logs to
  rotate. Can be set in parallel with RotateSizeMB.
  By default this parameter is set to "1440".
  
  

**Rotation/SizeMB** (default: 1024, unit: mb)

  This value defines the maximum file size in MB that triggers a file rotate.
  Files can get bigger than this size.
  By default this parameter is set to "1024".
  
  

**Rotation/Timestamp** (default: 2006-01-02_15)

  This value sets the timestamp added to the filename when file rotation
  is enabled. The format is based on Go's time.Format function.
  By default this parameter is to to "2006-01-02_15".
  
  

**Rotation/ZeroPadding** (default: 0)

  This value sets the number of leading zeros when rotating files with
  an existing name. Setting this setting to 0 won't add zeros, every other
  number defines the number of leading zeros to be used.
  By default this parameter is set to "0".
  
  

**Rotation/Compress** (default: false)

  This value defines if a rotated logfile is to be gzip compressed or not.
  By default this parameter is set to "false".
  
  

**Rotation/At**

  This value defines a specific time for rotation in hh:mm format.
  By default this parameter is set to "".
  
  

**Rotation/AtHour** (default: -1)

  (no documentation available)
  

**Rotation/AtMin** (default: -1)

  (no documentation available)
  

Parameters (from core.BufferedProducer)
---------------------------------------

**Channel**

  This value defines the capacity of the message buffer.
  By default this parameter is set to "8192".
  
  

**ChannelTimeoutMs** (default: 0, unit: ms)

  This value defines a timeout for each message
  before the message will discarded. To disable the timeout, set this
  parameter to 0.
  By default this parameter is set to "0".
  
  

Parameters (from core.SimpleProducer)
-------------------------------------

**Streams**

  Defines a list of streams the producer will receive from. This
  parameter is mandatory. Specifying "*" causes the producer to receive messages
  from all streams except internal internal ones (e.g. _GOLLUM_).
  By default this parameter is set to an empty list.
  
  

**FallbackStream**

  Defines a stream to route messages to if delivery fails.
  The message is reset to its original state before being routed, i.e. all
  modifications done to the message after leaving the consumer are removed.
  Setting this paramater to "" will cause messages to be discared when delivery
  fails.
  
  

**ShutdownTimeoutMs** (default: 1000, unit: ms)

  Defines the maximum time in milliseconds a producer is
  allowed to take to shut down. After this timeout the producer is always
  considered to have shut down.  Decreasing this value may lead to lost
  messages during shutdown. Raising it may increase shutdown time.
  
  

**Modulators**

  Defines a list of modulators to be applied to a message when
  it arrives at this producer. If a modulator changes the stream of a message
  the message is NOT routed to this stream anymore.
  By default this parameter is set to an empty list.
  
  

Examples
--------

This example will collect messages from the fallback stream and buffer them
for 10 minutes. After 10 minutes the first messages will be written back to
the system as fast as possible.

.. code-block:: yaml

	 spooling:
	   Type: producer.Spooling
	   Stream: fallback
	   MaxMessagesSec: 0
	   MaxFileAgeMin: 10





