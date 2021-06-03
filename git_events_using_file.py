#!/usr/bin/env python2
import re
import csv
import json

all_user_events = open("events.csv","r").read()

def write_and_update_file(event_name, event_type):
    row_list = sorted(event_type.items() , key=lambda t : t[1], reverse=True)

    #print event_name
    with open(event_name, 'w+') as file:
        writer = csv.writer(file)
        writer.writerows(row_list)

def get_push_events(actor_id=True):
    push_events = {}
    event_file = "push_event.csv"
    matches = re.findall(r'.*PushEvent.*', all_user_events)

    top_push_evets= {}

    for match in matches:
        events =  match.split(",")
        if actor_id:
            id = events[2]
        else:
            id = events[3]

        if id not in push_events:
            push_events[id] = 1
        else:
            push_events[id] += 1


    write_and_update_file(event_file, push_events)

    #read top 50 push events
    count = 0
    with open(event_file, 'r') as fd:
        events = csv.reader(fd)
        for event in events:
            top_push_evets[event[0]] = event[1]
            if count == 50:
                break

            count +=1

    return top_push_evets

def get_pull_events():
    pull_events = {}
    matches = re.findall(r'.*PullRequestEvent.*', all_user_events)

    for match in matches:
        events =  match.split(",")
        actor_id = events[2]
        if actor_id not in pull_events:
            pull_events[actor_id] = 1
        else:
            pull_events[actor_id] += 1

    return pull_events

def get_watch_events():
    watch_events = {}
    matches = re.findall(r'.*WatchEvent.*', all_user_events)

    for match in matches:
        events =  match.split(",")
        repo_id = events[3]
        if repo_id not in watch_events:
            watch_events[repo_id] = 1
        else:
            watch_events[repo_id] += 1

    result = sorted(watch_events.items() , key=lambda t : t[1], reverse=True)
    return result

def top10_active_users():
    all_actors = open("actors.csv","r").read()
    all_push_event = get_push_events()

    all_commit_count = all_push_event
    all_pull_event = get_pull_events()
    user_events = {}
    for user_id, push_event_count in all_push_event.items():
        if user_id in all_pull_event:
            user_events[user_id] = int(push_event_count) + all_pull_event[user_id]
        else:
            user_events[user_id] = int(push_event_count)

    result = sorted(user_events.items() , key=lambda t : t[1], reverse=True)
    print "username : PR_count : commits_count"
    print "-----------------------------------"
    for res in result[0:10]:
        user_id = res[0]
        commit_num = all_commit_count.get(user_id, 0)
        pr_num = all_pull_event.get(user_id, 0)
        user = user_id+".*"
        match = re.search(user, all_actors)
        user_name = match.group().split(",")[1]
        print user_name,":",pr_num,":",commit_num

def top10_repo_commits():
    all_repos = open("repos.csv","r").read()
    event_file = "push_event.csv"
    get_push_events(False)

    print "repo_name : commits_count"
    print "-------------------------"
    c = 0
    with open(event_file, 'r') as fd:
        events = csv.reader(fd)
        for event in events:
            repo_id = event[0]
            count = int(event[1])
            repo_name = repo_id+".*"

            match = re.search(repo_name, all_repos)
            reponame = match.group().split(",")[1]
            print reponame,":", count

            if c == 10:
                break

            c +=1

def top10_repo_watch_event():
    all_repos = open("repos.csv","r").read()
    watch_result = get_watch_events()

    print "repo_name : watch_count"
    print "-------------------------"
    for result in watch_result[0:10]:
        repo_id = result[0]
        repo_name = repo_id+".*"
        match = re.search(repo_name, all_repos)
        reponame = match.group().split(",")[1]
        print repo_id, reponame,":", result[1]

if __name__ == "__main__":
    print "------------------------------------------------------------------------------------------------------"
    print "Top 10 active users by amount of PRs created and commits pushed (username, PRs count, commits count) \n"
    top10_active_users()
    print "------------------------------------------------------------------------------------------------------\n"

    print "------------------------------------------------------------------------------------------------------"
    print "Top 10 repositories by amount of commits pushed (repo name, commits count) \n"
    top10_repo_commits()
    print "------------------------------------------------------------------------------------------------------\n"

    print "------------------------------------------------------------------------------------------------------"
    print "Top 10 repositories by amount of watch events (repo name, watch events count) \n"
    top10_repo_watch_event()
    print "------------------------------------------------------------------------------------------------------"
    
