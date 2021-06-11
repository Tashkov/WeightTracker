# WeightTracker
Use Cases:
/user
  1. Create user;
    -first_name
    -last_name
    -age
    -height
    -sex
    -weight_current
    -weight_goal
    -desired_weight_loss (weakly/monthly?)
  2. Read user info:
    1. Display all user in DB.
    2. Display specific users information:
      3. Display specific information about the user.
  3. Update user information;
  4. Delete user information.
  5. Display All Users // Completed 

/weight_logs
  1. Create weight log;
    -link to single user
    -date of log
  2. Read specific users weight logs:
  3. Update weight log
  4. Delete weight log

/app_functionality
  1. Calculate average weight;
    1. Calculate average weight for a specific time period.
  2. Alerts:  
    1. If average weight is going up over a fixed time period.
    2. If weight loss is progressing faster than the desired_weight_loss
    3. If weight gain instead of weight loss.
  3. Motivational notifications:
    1. For consecutive weight logs (weekly/monthly)
    2. For loosing the desired_weight_loss
    
