<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: helloworld.proto

namespace Helloworld;

use Google\Protobuf\Internal\DescriptorPool;
use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

class HelloRequest extends \Google\Protobuf\Internal\Message
{
    private $name = '';

    public function getName()
    {
        return $this->name;
    }

    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;
    }

}

class HelloReply extends \Google\Protobuf\Internal\Message
{
    private $message = '';

    public function getMessage()
    {
        return $this->message;
    }

    public function setMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->message = $var;
    }

}

$pool = DescriptorPool::getGeneratedPool();

$pool->internalAddGeneratedFile(hex2bin(
    "0ae6010a1068656c6c6f776f726c642e70726f746f120a68656c6c6f776f" .
    "726c64221c0a0c48656c6c6f52657175657374120c0a046e616d65180120" .
    "012809221d0a0a48656c6c6f5265706c79120f0a076d6573736167651801" .
    "2001280932490a0747726565746572123e0a0853617948656c6c6f12182e" .
    "68656c6c6f776f726c642e48656c6c6f526571756573741a162e68656c6c" .
    "6f776f726c642e48656c6c6f5265706c79220042360a1b696f2e67727063" .
    "2e6578616d706c65732e68656c6c6f776f726c64420f48656c6c6f576f72" .
    "6c6450726f746f5001a20203484c57620670726f746f33"
));
