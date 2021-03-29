import cv2
import math

img = cv2.imread(
    '/Users/wangyang/Downloads/8dbf5e27-243b-4bd9-866d-5ada1e4237be.jpg')

# img = cv2.imread(
#     '/Users/wangyang/Downloads/source.jpeg')

imgray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

ret, thresh = cv2.threshold(imgray, 200, 255, 0)
# cv2.imshow('src', thresh)
# cv2.waitKey()

# cv2.Canny

contours, hierarchy = cv2.findContours(
    thresh, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)

print("contours:", len(contours))


# cnt = contours[0]
# M = cv2.boundingRect(cnt)  # 以dict形式输出所有moment值
# print(M)

contours.sort(key=cv2.contourArea, reverse=True)

showCountour = []

newImg = img.copy()


ary = []
for index in range(len(contours)):
    contour = contours[index]
    x, y, w, h = cv2.boundingRect(contour)
    if (w > 500 and h > 500) or (w < 100 and h < 100):
        continue
    ary.append(w*h)
    # print('i:{0}\tw:{1}\th:{2}\tarea:{3}'.format(index, w, h, w*h))
    showCountour.append(contour)
    cv2.rectangle(newImg, (x, y), (x+w, y+h), (0, 0, 255), 1)
    cv2.putText(newImg, 'i:{0},w:{1},h:{2}'.format(index, w, h), (x, y),
                cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 255, 0), 1)

cv2.imwrite("/Users/wangyang/Downloads/new.jpg", newImg)

newImg = cv2.drawContours(img, showCountour, -1, (0, 0, 255), 1)
cv2.imwrite("/Users/wangyang/Downloads/new-s.jpg", newImg)
# cv2.imshow('src', img)

# cv2.waitKey()

ary.sort()

print(len(ary), ary)


dist = {1: 2}
str(dist)


for index in range(len(ary)):
    ceilNum = math.ceil(ary[index]/1000)
    if(ceilNum not in dist):
        dist[ceilNum] = 1
    else:
        dist[ceilNum] = dist[ceilNum]+1

for k, v in dist.items():
    print("area:{},qty:{},".format(k, v))
    # if(v < 4):
    #     continue
    # print(" type: '{}~{}', value: {}".format(k*1000, (k+1)*1000, v))

    # size = 100
    # minNum = ary[0]
    # maxNum = ary[1]

    # sizeSpan = ary[len(ary)-1]-ary[0]
    # for index in range(len(ary)):
    #     if(index+100) >= len(ary):
    #         break
    #     currentNumber = ary[index]
    #     maxNumber = ary[index+100]
    #     currentSpan = maxNumber-currentNumber
    #     if(currentSpan < sizeSpan):
    #         sizeSpan = currentSpan
    #         minNum = currentNumber
    #         maxNum = maxNumber

    # print(minNum, maxNum)
