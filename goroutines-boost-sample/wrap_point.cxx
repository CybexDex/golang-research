#include "wrap_point.hxx"
#include "point.hxx"

#include <boost/thread.hpp>
#include <boost/chrono.hpp>
#include <iostream>

void wait(int seconds)
{
  boost::this_thread::sleep_for(boost::chrono::seconds{seconds});
}

void thread()
{
  try
  {
    for (int i = 0; i < 5; ++i)
    {
      wait(1);
      auto res = Point(i, i).distance_to(Point(i, i));
      std::cout << res << '\n';
    }
  }
  catch (boost::thread_interrupted&) {}
}




double distance_between(double x1, double y1, double x2, double y2) {
  boost::thread t{thread};
  wait(3);
  t.interrupt();
  t.join();
}
