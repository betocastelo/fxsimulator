# Summary

Simple foreign exchange forward simulation, which assumes that the exchange rate follows a simple log-normal process
 with constant volatility. Uses the formula F(T) = F0\*e^-(s\*X\*T+1/2\*s^2*T), where
 * F0 is the expected forward rate,
 * s is the market volatility,
 * T is the time until maturity, and
 * X is a normal random variable with zero mean and variance 1.

This was a programming exercise during the interview process for my current job. I took the point of the test to be to
 assess my coding style, testing approach, and general form.