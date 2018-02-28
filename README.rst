
|fsutil| |travis| 

.. |godoc| image:: https://godoc.org/github.com/ymgyt/fsutil?status.svg
    :target: https://godoc.org/github.com/ymgyt/fsutil

.. |travis| image:: https://travis-ci.org/ymgyt/fsutil.svg?branch=master
    :target: https://travis-ci.org/ymgyt/fsutil


======
fsutil
======

fsutil provides file system utitily functions


Usage
=====

don't leak closing error
________________________

.. code-block:: go

   // do not this
   f, err := os.Open("test.txt")
   if err != nil {
    panic(err)
   }
   defer f.Close()

   // instead of simply defering Close()
   f, err := os.Open("test.txt")
   if err != nil {
    panic(err)
   }
   fsutil.Close(f, &err)
   return err


make sure Sync when writing
---------------------------

.. code-block:: go

   f, err := os.Create("content.txt")
   if err != nil {
    panic(err)
   }
   f.WriteString("content")
   defer f.Close()
   return f.Sync()

   // or
   f, err := os.Create("content.txt")
   if err != nil {
    panic(err)
   }
   f.WriteString("content")
   defer fsutil.SyncClose(f, &err)
   return err


