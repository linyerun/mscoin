#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename usercenter.json" -api usercenter.api -dir .

